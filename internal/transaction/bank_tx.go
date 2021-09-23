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

func BankSend(keyName string, sendOpt BankSendOption) error {
	// build tx context
	clientCtx := client.Context{}
	SetContextFromKeyName(clientCtx, keyName)
	txf := NewFactoryCLI(clientCtx)

	// build msg for tx
	toAddr := sendOpt.ToAddr
	fromAddr := clientCtx.GetFromAddress()
	coin := sdk.Coin{Denom: sendOpt.Denom, Amount: sendOpt.Amount}
	coins := sdk.Coins([]sdk.Coin{coin})
	msg := types.NewMsgSend(fromAddr, toAddr, coins)

	return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
}
