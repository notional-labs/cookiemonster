package transaction

import (
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

type BankSendOption struct {
	ToAddr sdk.AccAddress
	Denom  string
	Amount sdk.Int
}

func BankSend(txOpt TxOption, sendOpt BankSendOption) error {
	// build msg for tx
	toAddr := sendOpt.ToAddr
	fromAddr := txOpt.FromAddr
	coin := sdk.Coin{Denom: sendOpt.Denom, Amount: sendOpt.Amount}
	coins := sdk.Coins([]sdk.Coin{coin})
	msg := types.NewMsgSend(fromAddr, toAddr, coins)

	// build tx context
	clientCtx := client.Context{}
	SetContextFromTxOption(clientCtx, txOpt)
	txf := NewFactoryCLI(clientCtx)
	return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
}
