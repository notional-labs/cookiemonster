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
	code, log, txHash, err := BroadcastTx(clientCtx, txf, msg)
	if err != nil {
		return txHash, err
	}
	if code != 0 {
		return txHash, fmt.Errorf("tx failed with code %d with log = %s", code, log)
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
	Hash        string
}

func (joinPoolTx JoinPoolTx) Execute() (string, error) {

	keyName := joinPoolTx.KeyName
	joinPoolOpt := joinPoolTx.JoinPoolOpt

	gas := 200000
	var err error
	var txHash string

	// if tx failed because of insufficient fee , retry
	for i := 0; i < 4; i++ {
		fmt.Println("\n---------------")
		fmt.Printf("\n Try %d times\n\n", i+1)
		txHash, err = JoinPool(keyName, joinPoolOpt, uint64(gas))
		if err == nil {
			joinPoolTx.Hash = txHash
			return txHash, nil
		}
		if err.Error() == "insufficient fee" {
			fmt.Println("\nTx failed because of insufficient fee, try again with higher gas")
			gas += 300000
		} else {
			fmt.Println("\n" + err.Error() + " try again")
		}
	}
	return txHash, err
}

func (joinPoolTx JoinPoolTx) Report(reportPath string) {

	joinPoolOpt := joinPoolTx.JoinPoolOpt
	keyName := joinPoolTx.KeyName

	f, _ := os.OpenFile(reportPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	n, err := f.WriteString("\nJoin Pool Transaction\n")
	if err != nil {
		fmt.Println(n, err)
	}
	n, err = f.WriteString("\nKeyname: " + keyName + "\n")
	if err != nil {
		fmt.Println(n, err)
	}
	n, err = f.WriteString("\nJoin Pool Option\n\n")
	if err != nil {
		fmt.Println(n, err)
	}

	txData, _ := yaml.Marshal(joinPoolOpt)
	_, _ = f.Write(txData)
	n, err = f.WriteString("\ntx hash: " + joinPoolTx.Hash + "\n")
	if err != nil {
		fmt.Println(n, err)
	}
	n, err = f.WriteString(Seperator)
	if err != nil {
		fmt.Println(n, err)
	}

	f.Close()
}

func (joinPoolTx JoinPoolTx) Prompt() {
	joinPoolOpt := joinPoolTx.JoinPoolOpt
	keyName := joinPoolTx.KeyName
	fmt.Print(Seperator)
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

	code, log, txHash, err := BroadcastTx(clientCtx, txf, msg)
	if err != nil {
		return txHash, err
	}
	if code == 11 {
		return txHash, fmt.Errorf("insufficient fee")
	}
	if code != 0 {
		return txHash, fmt.Errorf("tx failed with code %d with log = %s", code, log)
	}

	fmt.Println("hello1")

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
	Hash           string
}

func (swapAndPoolTx SwapAndPoolTx) Execute() (string, error) {

	keyName := swapAndPoolTx.KeyName
	swapAndPoolOpt := swapAndPoolTx.SwapAndPoolOpt
	gas := 1000000
	var err error
	var txHash string

	// if tx failed because of insufficient fee , retry
	for i := 0; i < 4; i++ {
		fmt.Println("\n---------------")
		fmt.Printf("\n Try %d times\n\n", i+1)
		txHash, err = SwapAndPool(keyName, swapAndPoolOpt, uint64(gas))
		if err == nil {
			swapAndPoolTx.Hash = txHash
			return txHash, nil
		}
		if err.Error() == "insufficient fee" {
			fmt.Println("\nTx failed because of insufficient fee, try again with higher gas")
			gas += 300000
		} else {
			fmt.Println("\n" + err.Error() + " try again")
		}
	}
	return txHash, err
}

func (swapAndPoolTx SwapAndPoolTx) Report(reportPath string) {

	swapAndPoolOpt := swapAndPoolTx.SwapAndPoolOpt
	keyName := swapAndPoolTx.KeyName

	f, _ := os.OpenFile(reportPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	n, err := f.WriteString("\nSwap And Pool Transaction\n")
	if err != nil {
		fmt.Println(err, n)
	}
	n, err = f.WriteString("\nKeyname: " + keyName + "\n")
	if err != nil {
		fmt.Println(err, n)
	}

	n, err = f.WriteString("\nSwap And Pool Option\n\n")
	if err != nil {
		fmt.Println(err, n)
	}

	txData, _ := yaml.Marshal(swapAndPoolOpt)
	_, _ = f.Write(txData)
	n, err = f.WriteString("\ntx hash: " + swapAndPoolTx.Hash + "\n")
	if err != nil {
		fmt.Println(err, n)
	}
	n, err = f.WriteString(Seperator)
	if err != nil {
		fmt.Println(err, n)
	}

	f.Close()
}

func (swapAndPoolTx SwapAndPoolTx) Prompt() {
	swapAndPoolOpt := swapAndPoolTx.SwapAndPoolOpt
	keyName := swapAndPoolTx.KeyName
	fmt.Print(Seperator)
	fmt.Print("\nSwap And Pool Transaction\n")
	fmt.Print("\nKeyname: " + keyName + "\n")
	fmt.Print("\nSwap And Pool Option\n\n")
	fmt.Printf("%+v\n", swapAndPoolOpt)

}
