package transaction

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/notional-labs/cookiemonster/osmosis"
	"github.com/notional-labs/cookiemonster/query"
)

func ClaimReward(keyName string, gas uint64) (string, error) {
	clientCtx := osmosis.GetDefaultClientContext()
	clientCtx, err := SetKeyNameToContext(clientCtx, keyName)
	if err != nil {
		return "", err
	}
	delAddr := clientCtx.GetFromAddress()

	queryClient := types.NewQueryClient(clientCtx)
	delValsRes, err := queryClient.DelegatorValidators(context.Background(), &types.QueryDelegatorValidatorsRequest{DelegatorAddress: delAddr.String()})
	fmt.Println(delValsRes)
	if err != nil {
		return "", err
	}
	validators := delValsRes.Validators
	// build multi-message transaction
	msgs := make([]sdk.Msg, 0, len(validators))
	for _, valAddr := range validators {
		val, err := sdk.ValAddressFromBech32(valAddr)
		if err != nil {
			return "", err
		}

		msg := types.NewMsgWithdrawDelegatorReward(delAddr, val)
		if err := msg.ValidateBasic(); err != nil {
			return "", err
		}
		msgs = append(msgs, msg)
	}
	txf := NewTxFactoryFromClientCtx(clientCtx).WithGas(gas)
	code, txHash, err := BroadcastTx(clientCtx, txf, msgs...)
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

type ClaimTx struct {
	KeyName string
}

func (claimTx ClaimTx) Execute() (string, error) {
	keyName := claimTx.KeyName
	gas := 200000
	var err error
	var txHash string

	// if tx failed because of insufficient fee , retry
	for i := 0; i < 4; i++ {
		txHash, err = ClaimReward(keyName, uint64(gas))
		if err == nil {
			return txHash, nil
		}
		if err.Error() == "insufficient fee" {
			gas += 300000
		}
	}
	return txHash, err
}

func (claimTx ClaimTx) Report() {

	keyName := claimTx.KeyName

	f, _ := os.OpenFile("report", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	f.WriteString("\nClaim Reward Transaction\n")
	f.WriteString("\nKeyname: " + keyName + "\n")
	f.WriteString(transactionSeperator)

	f.Close()
}

func (claimTx ClaimTx) Prompt() {
	keyName := claimTx.KeyName
	fmt.Print(transactionSeperator)
	fmt.Print(transactionSeperator)
	fmt.Print("\nClaim Reward Transaction\n")
	fmt.Print("\nKeyname: " + keyName + "\n")
	fmt.Print("\nClaim Reward Option\n\n")

}
