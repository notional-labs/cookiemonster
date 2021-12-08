package transaction

import (
	"fmt"
	"os"
	"time"

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
	clientCtx := osmosis.GetDefaultClientContext()
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

	code, log, txHash, err := BroadcastTx(clientCtx, txf, msg)
	if err != nil {
		return txHash, err
	}
	if code != 0 {
		return txHash, fmt.Errorf("tx failed with code %d with log = %s", code, log)
	}

	broadcastedTx, err := query.QueryTxWithRetry(txHash, 4)
	if err != nil {
		return "", err
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
	gas := 2000000
	var err error
	var txHash string

	// if tx failed because of insufficient fee , retry
	for i := 0; i < 4; i++ {
		fmt.Println("\n---------------")
		fmt.Printf("\n Try %d times\n\n", i+1)
		txHash, err = BankSend(keyName, bankSendOpt, uint64(gas))

		if err == nil {
			bankSendTx.Hash = txHash
			return txHash, nil
		}
		if err.Error() == "insufficient fee" {
			fmt.Println("\nTx failed because of insufficient fee, try again with higher gas\n")
			gas += 300000
		} else {
			time.Sleep(5 * time.Second)

			fmt.Println("\n" + err.Error() + " try again\n")
		}
	}
	return txHash, err
}

func (bankSendTx BankSendTx) Report(reportPath string) {

	bankSendOpt := bankSendTx.BankSendOpt
	keyName := bankSendTx.KeyName
	hash := bankSendTx.Hash

	f, _ := os.OpenFile(reportPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	num, err := f.WriteString("\nBank Send Transaction\n")
	if err != nil {
		fmt.Println(err, "something this number is here", num)
	}
	num, err = f.WriteString("\nKeyname: " + keyName + "\n")
	if err != nil {
		fmt.Println(err, "something this number is here", num)
	}
	num, err = f.WriteString("\nBank Send Option\n\n")
	if err != nil {
		fmt.Println(err, "something this number is here", num)
	}

	txData, _ := yaml.Marshal(bankSendOpt)
	_, _ = f.Write(txData)
	num, err = f.WriteString("\ntx hash: " + hash + "\n")
	if err != nil {
		fmt.Println(num, err)
	}
	f.WriteString(Seperator)

	f.Close()
}

func (bankSendTx BankSendTx) Prompt() {
	bankSendOpt := bankSendTx.BankSendOpt
	keyName := bankSendTx.KeyName
	fmt.Print(Seperator)

	fmt.Print("\nBank Send Transaction\n")
	fmt.Print("\nKeyname: " + keyName + "\n")
	fmt.Print("\nBank Send Option\n\n")
	fmt.Printf("%+v\n", bankSendOpt)
}
