package transaction

import (
	"context"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/notional-labs/cookiemonster/osmosis"
)

func ClaimReward(keyName string) error {
	clientCtx := osmosis.DefaultClientCtx
	clientCtx, err := SetKeyNameToContext(clientCtx, keyName)
	if err != nil {
		return err
	}
	delAddr := clientCtx.GetFromAddress()

	// The transaction cannot be generated offline since it requires a query
	// to get all the validators.
	if clientCtx.Offline {
		return fmt.Errorf("cannot generate tx in offline mode")
	}

	queryClient := types.NewQueryClient(clientCtx)
	delValsRes, err := queryClient.DelegatorValidators(context.Background(), &types.QueryDelegatorValidatorsRequest{DelegatorAddress: delAddr.String()})
	if err != nil {
		return err
	}

	validators := delValsRes.Validators
	// build multi-message transaction
	msgs := make([]sdk.Msg, 0, len(validators))
	for _, valAddr := range validators {
		val, err := sdk.ValAddressFromBech32(valAddr)
		if err != nil {
			return err
		}

		msg := types.NewMsgWithdrawDelegatorReward(delAddr, val)
		if err := msg.ValidateBasic(); err != nil {
			return err
		}
		msgs = append(msgs, msg)
	}

	chunkSize := 0
	if clientCtx.BroadcastMode != flags.BroadcastBlock && chunkSize > 0 {
		return fmt.Errorf("cannot use broadcast mode %[1]s with max-msgs != 0",
			clientCtx.BroadcastMode)
	}

	return newSplitAndApply(clientCtx, msgs, chunkSize)
}

func newSplitAndApply(clientCtx client.Context, msgs []sdk.Msg, chunkSize int,
) error {

	if chunkSize == 0 {
		txf := NewTxFactoryFromClientCtx(clientCtx)
		return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msgs...)
	}

	// split messages into slices of length chunkSize
	totalMessages := len(msgs)
	for i := 0; i < len(msgs); i += chunkSize {

		sliceEnd := i + chunkSize
		if sliceEnd > totalMessages {
			sliceEnd = totalMessages
		}

		msgChunk := msgs[i:sliceEnd]
		txf := NewTxFactoryFromClientCtx(clientCtx)
		if err := tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msgChunk...); err != nil {
			return err
		}
	}
	return nil
}

type ClaimTx struct {
	KeyName string
}

func (claimTx ClaimTx) Execute() error {
	keyName := claimTx.KeyName
	err := ClaimReward(keyName)
	return err
}

func (claimTx ClaimTx) Report() {

	keyName := claimTx.KeyName

	f, _ := os.OpenFile("report", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	f.WriteString("\nClaim Reward Transaction\n")
	f.WriteString("\nKeyname: " + keyName + "\n")
	f.WriteString(transactionSeperator)

	f.Close()
}
