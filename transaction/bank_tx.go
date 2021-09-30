package transaction

import (
	"fmt"
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/notional-labs/cookiemonster/osmosis"
	"github.com/notional-labs/cookiemonster/query"
	"gopkg.in/yaml.v3"
)

type BankSendOption struct {
	ToAddr sdk.AccAddress
	Denom  string
	Amount sdk.Int
}

func BankSend(keyName string, bankSendOpt BankSendOption, gas uint64) (string, error) {
	// build tx context
	clientCtx := osmosis.DefaultClientCtx
	clientCtx, err := SetKeyNameToContext(clientCtx, keyName)
	if err != nil {
		return "", err
	}
	txf := NewTxFactoryFromClientCtx(clientCtx).WithGas(gas)

	// build msg for tx
	toAddr := bankSendOpt.ToAddr
	fromAddr := clientCtx.GetFromAddress()
	coin := sdk.Coin{Denom: bankSendOpt.Denom, Amount: bankSendOpt.Amount}
	coins := sdk.Coins([]sdk.Coin{coin})
	msg := types.NewMsgSend(fromAddr, toAddr, coins)

	code, txHash, err := BroadcastTx(clientCtx, txf, msg)
	if err != nil {
		return txHash, err
	}
	if code != 0 {
		return txHash, fmt.Errorf("tx failed with code %d", code)
	}
	broadcastedTx, err := query.QueryTx(txHash)
	if err != nil {
		return txHash, err
	}
	if broadcastedTx.Code == 11 {
		return txHash, fmt.Errorf("insufficient fee")

	}
	if broadcastedTx.Code != 0 {
		return txHash, fmt.Errorf("tx failed with code %d", code)
	}
	return txHash, nil
}

type BankSendTx struct {
	BankSendOpt BankSendOption
	KeyName     string
	Hash        string
}

func (bankSendTx BankSendTx) Execute() (string, error) {
	keyName := bankSendTx.KeyName
	bankSendOpt := bankSendTx.BankSendOpt
	gas := 200000
	var err error
	var txHash string

	// if tx failed because of insufficient fee , retry
	for i := 0; i < 4; i++ {
		txHash, err = BankSend(keyName, bankSendOpt, uint64(gas))
		bankSendTx.Hash = txHash
		if err == nil {
			return txHash, nil
		}
		if err.Error() != "insufficient fee" {
			return txHash, err
		}
		gas += 300000
	}
	return txHash, err
}

func (bankSendTx BankSendTx) Report() {

	bankSendOpt := bankSendTx.BankSendOpt
	keyName := bankSendTx.KeyName
	hash := bankSendTx.Hash

	f, _ := os.OpenFile("report", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	f.WriteString("\nBank Send Transaction\n")
	f.WriteString("\nKeyname: " + keyName + "\n")
	f.WriteString("\nBank Send Option\n\n")

	txData, _ := yaml.Marshal(bankSendOpt)
	_, _ = f.Write(txData)
	f.WriteString("\ntx hash: " + hash + "\n")
	f.WriteString(transactionSeperator)

	f.Close()
}

func (bankSendTx BankSendTx) Prompt() {
	bankSendOpt := bankSendTx.BankSendOpt
	keyName := bankSendTx.KeyName
	fmt.Print(transactionSeperator)

	fmt.Print("\nBank Send Transaction\n")
	fmt.Print("\nKeyname: " + keyName + "\n")
	fmt.Print("\nBank Send Option\n\n")
	fmt.Printf("%+v\n", bankSendOpt)
}
