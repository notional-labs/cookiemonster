package transaction

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
)

func ClaimReward(txOpt TxOption) error {
	clientCtx := client.Context{}
	SetContextFromTxOption(clientCtx, txOpt)

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
		txf := NewFactoryCLI(clientCtx)
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
		txf := NewFactoryCLI(clientCtx)
		if err := tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msgChunk...); err != nil {
			return err
		}
	}
	return nil
}
