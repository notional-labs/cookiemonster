package transaction

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/osmosis-labs/osmosis/x/gamm/types"
)

type PoolOption struct {
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

func JoinPool(keyName string, poolOpt PoolOption) error {
	// build tx context
	clientCtx := client.Context{}
	SetContextFromKeyName(clientCtx, keyName)
	txf := NewFactoryCLI(clientCtx)

	// build msg for tx
	fromAddr := clientCtx.FromAddress
	poolId := poolOpt.PoolId
	shareOutAmount := poolOpt.ShareOutAmount
	maxAmountsIn := poolOpt.MaxAmountsIn

	msg := NewMsgJoinPool(fromAddr, poolId, shareOutAmount, maxAmountsIn)

	return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
}
