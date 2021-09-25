package transaction

import (
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/notional-labs/cookiemonster/internal/osmosis"
)

type DelegateOption struct {
	ValAddr sdk.ValAddress
	Denom   string
	Amount  sdk.Int
}

func Delegate(keyName string, delOpt DelegateOption) error {
	// build tx context
	clientCtx := osmosis.DefaultClientCtx
	clientCtx, err := ContextWithKeyName(clientCtx, keyName)
	if err != nil {
		return err
	}

	txf := NewFactoryCLI(clientCtx)

	// build msg for tx
	valAddr := delOpt.ValAddr
	delAddr := clientCtx.FromAddress
	amount := sdk.Coin{Denom: delOpt.Denom, Amount: delOpt.Amount}
	msg := types.NewMsgDelegate(delAddr, valAddr, amount)

	return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
}

type UndelegateOption struct {
	ValAddr sdk.ValAddress
	Denom   string
	Amount  sdk.Int
}

func Undelegate(keyName string, undelOpt UndelegateOption) error {
	// build tx context
	clientCtx := osmosis.DefaultClientCtx
	clientCtx, err := ContextWithKeyName(clientCtx, keyName)
	if err != nil {
		return err
	}

	txf := NewFactoryCLI(clientCtx)

	// build msg for tx
	valAddr := undelOpt.ValAddr
	delAddr := clientCtx.FromAddress
	amount := sdk.Coin{Denom: undelOpt.Denom, Amount: undelOpt.Amount}
	msg := types.NewMsgDelegate(delAddr, valAddr, amount)

	return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
}
