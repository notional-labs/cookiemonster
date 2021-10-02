package transaction

import (
	"fmt"
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/cookiemonster/osmosis"
	"github.com/notional-labs/cookiemonster/query"
	"github.com/osmosis-labs/osmosis/x/gamm/types"
	"gopkg.in/yaml.v3"
)

type JoinPoolOption struct {
	PoolId         uint64
	ShareOutAmount sdk.Int
	MaxAmountsIn   sdk.Coins
}

func NewMsgJoinPool(fromAddr sdk.AccAddress, poolId uint64, shareOutAmount sdk.Int, maxAmountsIn sdk.Coins) sdk.Msg {
	msg := &types.MsgJoinPool{
		Sender:         fromAddr.String(),
		PoolId:         poolId,
		ShareOutAmount: shareOutAmount,
		TokenInMaxs:    maxAmountsIn,
	}
	return msg
}

func JoinPool(keyName string, joinPoolOpt JoinPoolOption, gas uint64) (string, error) {
	// build tx context
	clientCtx := osmosis.GetDefaultClientContext()
	clientCtx, err := SetKeyNameToContext(clientCtx, keyName)
	if err != nil {
		return "", err
	}
	txf := NewTxFactoryFromClientCtx(clientCtx).WithGas(gas)

	// build msg for tx
	fromAddr := clientCtx.FromAddress
	poolId := joinPoolOpt.PoolId
	shareOutAmount := joinPoolOpt.ShareOutAmount
	maxAmountsIn := joinPoolOpt.MaxAmountsIn

	msg := NewMsgJoinPool(fromAddr, poolId, shareOutAmount, maxAmountsIn)
	code, txHash, err := BroadcastTx(clientCtx, txf, msg)
	if err != nil {
		return txHash, err
	}
	if code != 0 {
		return txHash, fmt.Errorf("tx failed with code %d", code)
	}
	broadcastedTx, err := query.QueryTxWithRetry(txHash, 4)
	if err != nil {
		return txHash, err
	}
	if broadcastedTx.Code == 11 {
		return txHash, fmt.Errorf("insufficient fee")

	}
	if broadcastedTx.Code != 0 {
		return txHash, fmt.Errorf("tx failed with code %d", code)
	}
	return txHash, nil
}

type JoinPoolTx struct {
	KeyName     string
	JoinPoolOpt JoinPoolOption
}

func (joinPoolTx JoinPoolTx) Execute() (string, error) {

	keyName := joinPoolTx.KeyName
	joinPoolOpt := joinPoolTx.JoinPoolOpt

	gas := 200000
	var err error
	var txHash string

	// if tx failed because of insufficient fee , retry
	for i := 0; i < 4; i++ {
		txHash, err = JoinPool(keyName, joinPoolOpt, uint64(gas))
		if err == nil {
			return txHash, nil
		}
		if err.Error() != "insufficient fee" {
			return txHash, err
		}
		gas += 300000
	}
	return txHash, err
}

func (joinPoolTx JoinPoolTx) Report() {

	joinPoolOpt := joinPoolTx.JoinPoolOpt
	keyName := joinPoolTx.KeyName

	f, _ := os.OpenFile("report", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	f.WriteString("\nJoin Pool Transaction\n")
	f.WriteString("\nKeyname: " + keyName + "\n")
	f.WriteString("\nJoin Pool Option\n\n")

	txData, _ := yaml.Marshal(joinPoolOpt)
	_, _ = f.Write(txData)
	f.WriteString(transactionSeperator)

	f.Close()
}

func (joinPoolTx JoinPoolTx) Prompt() {
	joinPoolOpt := joinPoolTx.JoinPoolOpt
	keyName := joinPoolTx.KeyName
	fmt.Print(transactionSeperator)
	fmt.Print("\nJoin Pool Transaction\n")
	fmt.Print("\nKeyname: " + keyName + "\n")
	fmt.Print("\nJoin Pool Option\n\n")
	fmt.Printf("%+v\n", joinPoolOpt)

}

type SwapAndPoolOption struct {
	PoolId            uint64
	TokenInAmount     sdk.Int
	TokenInDenom      string
	ShareOutMinAmount sdk.Int
}

func SwapAndPool(keyName string, swapAndPoolOption SwapAndPoolOption, gas uint64) (string, error) {
	// build tx context
	clientCtx := osmosis.GetDefaultClientContext()
	clientCtx, err := SetKeyNameToContext(clientCtx, keyName)
	if err != nil {
		return "", err
	}
	txf := NewTxFactoryFromClientCtx(clientCtx).WithGas(gas)

	// build msg for tx
	fromAddr := clientCtx.FromAddress
	poolId := swapAndPoolOption.PoolId
	shareOutMinAmount := swapAndPoolOption.ShareOutMinAmount
	tokenIn := sdk.Coin{Amount: swapAndPoolOption.TokenInAmount, Denom: swapAndPoolOption.TokenInDenom}

	msg := &types.MsgJoinSwapExternAmountIn{
		Sender:            fromAddr.String(),
		PoolId:            poolId,
		TokenIn:           tokenIn,
		ShareOutMinAmount: shareOutMinAmount,
	}
	code, txHash, err := BroadcastTx(clientCtx, txf, msg)
	if err != nil {
		return txHash, err
	}
	if code == 11 {
		return txHash, fmt.Errorf("insufficient fee")
	}
	if code != 0 {
		return txHash, fmt.Errorf("tx failed with code %d", code)
	}
	broadcastedTx, err := query.QueryTxWithRetry(txHash, 4)
	if err != nil {
		return txHash, err
	}
	if broadcastedTx.Code == 11 {
		return txHash, fmt.Errorf("insufficient fee")

	}
	if broadcastedTx.Code != 0 {
		return txHash, fmt.Errorf("tx failed with code %d", code)
	}
	return txHash, nil
}

type SwapAndPoolTx struct {
	KeyName        string
	SwapAndPoolOpt SwapAndPoolOption
}

func (swapAndPoolTx SwapAndPoolTx) Execute() (string, error) {

	keyName := swapAndPoolTx.KeyName
	swapAndPoolOpt := swapAndPoolTx.SwapAndPoolOpt
	gas := 200000
	var err error
	var txHash string

	// if tx failed because of insufficient fee , retry
	for i := 0; i < 4; i++ {
		fmt.Println(i, "try")
		txHash, err = SwapAndPool(keyName, swapAndPoolOpt, uint64(gas))
		if err == nil {
			return txHash, nil
		}
		if err.Error() != "insufficient fee" {
			return txHash, err
		}
		gas += 300000
	}
	return txHash, err
}

func (swapAndPoolTx SwapAndPoolTx) Report() {

	swapAndPoolOpt := swapAndPoolTx.SwapAndPoolOpt
	keyName := swapAndPoolTx.KeyName

	f, _ := os.OpenFile("report", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	f.WriteString("\nSwap And Pool Transaction\n")
	f.WriteString("\nKeyname: " + keyName + "\n")
	f.WriteString("\nSwap And Pool Option\n\n")

	txData, _ := yaml.Marshal(swapAndPoolOpt)
	_, _ = f.Write(txData)
	f.WriteString(transactionSeperator)

	f.Close()
}

func (swapAndPoolTx SwapAndPoolTx) Prompt() {
	swapAndPoolOpt := swapAndPoolTx.SwapAndPoolOpt
	keyName := swapAndPoolTx.KeyName
	fmt.Print(transactionSeperator)
	fmt.Print("\nSwap And Pool Transaction\n")
	fmt.Print("\nKeyname: " + keyName + "\n")
	fmt.Print("\nSwap And Pool Option\n\n")
	fmt.Printf("%+v\n", swapAndPoolOpt)

}
