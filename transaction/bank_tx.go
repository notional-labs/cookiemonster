package transaction

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/notional-labs/cookiemonster/osmosis"
	"gopkg.in/yaml.v3"
)

type BankSendOption struct {
	ToAddr sdk.AccAddress
	Denom  string
	Amount sdk.Int
}

func BankSend(keyName string, bankSendOpt BankSendOption) error {
	// build tx context
	clientCtx := osmosis.DefaultClientCtx
	clientCtx, err := SetKeyNameToContext(clientCtx, keyName)
	if err != nil {
		return err
	}
	txf := NewTxFactoryFromClientCtx(clientCtx)

	// build msg for tx
	toAddr := bankSendOpt.ToAddr
	fromAddr := clientCtx.GetFromAddress()
	coin := sdk.Coin{Denom: bankSendOpt.Denom, Amount: bankSendOpt.Amount}
	coins := sdk.Coins([]sdk.Coin{coin})
	msg := types.NewMsgSend(fromAddr, toAddr, coins)

	return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
}

type BankSendTx struct {
	BankSendOpt BankSendOption
	KeyName     string
}

func (bankSendTx BankSendTx) Execute() error {
	keyName := bankSendTx.KeyName
	bankSendOpt := bankSendTx.BankSendOpt
	err := BankSend(keyName, bankSendOpt)
	return err
}

func (bankSendTx BankSendTx) Report() {

	bankSendOpt := bankSendTx.BankSendOpt
	keyName := bankSendTx.KeyName

	f, _ := os.OpenFile("report", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	f.WriteString("\nBank Send Transaction\n")
	f.WriteString("\nKeyname: " + keyName + "\n")
	f.WriteString("\nBank Send Option\n\n")

	txData, _ := yaml.Marshal(bankSendOpt)
	_, _ = f.Write(txData)
	f.WriteString(transactionSeperator)

	f.Close()
}
