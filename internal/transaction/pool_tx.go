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

func JoinPool(txOpt TxOption, poolOpt PoolOption) error {
	// build msg for tx
	fromAddr := txOpt.FromAddr
	poolId := poolOpt.PoolId
	shareOutAmount := poolOpt.ShareOutAmount
	maxAmountsIn := poolOpt.MaxAmountsIn

	msg := NewMsgJoinPool(fromAddr, poolId, shareOutAmount, maxAmountsIn)

	// build tx context
	clientCtx := client.Context{}
	SetContextFromTxOption(clientCtx, txOpt)
	txf := NewFactoryCLI(clientCtx)
	return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
}
