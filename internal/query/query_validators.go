package main

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	gocontext "context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/spf13/pflag"
)

type newGenerateOrBroadcastFunc func(client.Context, *pflag.FlagSet, ...sdk.Msg) error

func newSplitAndApply(
	genOrBroadcastFn newGenerateOrBroadcastFunc, clientCtx client.Context,
	fs *pflag.FlagSet, msgs []sdk.Msg, chunkSize int,
) error {

	if chunkSize == 0 {
		return genOrBroadcastFn(clientCtx, fs, msgs...)
	}

	// split messages into slices of length chunkSize
	totalMessages := len(msgs)
	for i := 0; i < len(msgs); i += chunkSize {

		sliceEnd := i + chunkSize
		if sliceEnd > totalMessages {
			sliceEnd = totalMessages
		}

		msgChunk := msgs[i:sliceEnd]
		if err := genOrBroadcastFn(clientCtx, fs, msgChunk...); err != nil {
			return err
		}
	}

	return nil
}


func withdrawFromDelegation(clientCtx client.Context, commandCtx gocontext.Context, flags *pflag.FlagSet) error {

	var delAddr sdk.AccAddress

	// The transaction cannot be generated offline since it requires a query
	// to get all the validators.
	if clientCtx.Offline {
		fmt.Errorf("cannot generate tx in offline mode")

	}

	queryClient := types.NewQueryClient(clientCtx)
	delValsRes, err := queryClient.DelegatorValidators(commandCtx, &types.QueryDelegatorValidatorsRequest{DelegatorAddress: clientCtx.FromAddress.String()})
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
		msgs = append(msgs, msg)
	}

	chunkSize :=  10   // cmd.Flags().GetInt(FlagMaxMessagesPerTx)
	return newSplitAndApply(tx.GenerateOrBroadcastTxCLI, clientCtx, flags, msgs, chunkSize)
}
