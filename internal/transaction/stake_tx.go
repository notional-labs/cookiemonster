package transaction

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

type DelegateOption struct {
	ValAddr sdk.ValAddress
	Denom   string
	Amount  sdk.Int
}

func Delegate(txOpt TxOption, delOpt DelegateOption) error {
	// build msg for tx
	valAddr := delOpt.ValAddr
	delAddr := txOpt.FromAddr
	amount := sdk.Coin{Denom: delOpt.Denom, Amount: delOpt.Amount}
	msg := types.NewMsgDelegate(delAddr, valAddr, amount)

	// build tx context
	clientCtx := client.Context{}
	setContextFromTxOption(clientCtx, txOpt)
	txf := NewFactoryCLI(clientCtx)
	return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
}

type UndelegateOption struct {
	ValAddr sdk.ValAddress
	Denom   string
	Amount  sdk.Int
}

func Undelegate(txOpt TxOption, undelOpt UndelegateOption) error {
	// build msg for tx
	valAddr := undelOpt.ValAddr
	delAddr := txOpt.FromAddr
	amount := sdk.Coin{Denom: undelOpt.Denom, Amount: undelOpt.Amount}
	msg := types.NewMsgDelegate(delAddr, valAddr, amount)

	// build tx context
	clientCtx := client.Context{}
	setContextFromTxOption(clientCtx, txOpt)
	txf := NewFactoryCLI(clientCtx)
	return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
}
