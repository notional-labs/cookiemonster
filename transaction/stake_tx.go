package transaction

import (
	"fmt"
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/notional-labs/cookiemonster/osmosis"
	"github.com/notional-labs/cookiemonster/query"
	"gopkg.in/yaml.v3"
)

type DelegateOption struct {
	ValAddr sdk.ValAddress
	Denom   string
	Amount  sdk.Int
}

func Delegate(keyName string, delegateOpt DelegateOption, gas uint64) (string, error) {
	// build tx context
	clientCtx := osmosis.GetDefaultClientContext()
	clientCtx, err := SetKeyNameToContext(clientCtx, keyName)
	if err != nil {
		return "", err
	}

	txf := NewTxFactoryFromClientCtx(clientCtx).WithGas(gas)

	// build msg for tx
	valAddr := delegateOpt.ValAddr
	delAddr := clientCtx.FromAddress
	amount := sdk.Coin{Denom: delegateOpt.Denom, Amount: delegateOpt.Amount}
	msg := types.NewMsgDelegate(delAddr, valAddr, amount)

	code, txHash, err := BroadcastTx(clientCtx, txf, msg)
	if err != nil {
		return txHash, err
	}
	if code != 0 {
		return txHash, fmt.Errorf("tx failed with code %d", code)
	}
	broadcastedTx, err := query.QueryTxWithRetry(txHash, 4)
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

type DelegateTx struct {
	KeyName     string
	DelegateOpt DelegateOption
	Hash        string
}

func (delegateTx DelegateTx) Execute() (string, error) {

	keyName := delegateTx.KeyName
	delegateOpt := delegateTx.DelegateOpt
	gas := 2000000
	var err error
	var txHash string

	// if tx failed because of insufficient fee , retry
	for i := 0; i < 4; i++ {
		fmt.Println("\n---------------")
		fmt.Printf("\n Try %d times\n\n", i+1)
		txHash, err = Delegate(keyName, delegateOpt, uint64(gas))
		if err == nil {
			delegateTx.Hash = txHash
			return txHash, nil
		}
		if err.Error() == "insufficient fee" {
			fmt.Println("\nTx failed because of insufficient fee, try again with higher gas\n")
			gas += 300000
		} else {
			fmt.Println("\n" + err.Error() + " try again\n")
		}
	}
	return txHash, err
}

func (delegateTx DelegateTx) Report(reportPath string) {

	delegateOpt := delegateTx.DelegateOpt
	keyName := delegateTx.KeyName

	f, _ := os.OpenFile(reportPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	f.WriteString("\nDelegate Transaction\n")
	f.WriteString("\nKeyname: " + keyName + "\n")
	f.WriteString("\nDelegate Option\n\n")

	txData, _ := yaml.Marshal(delegateOpt)
	_, _ = f.Write(txData)
	f.WriteString("\ntx hash: " + delegateTx.Hash + "\n")
	f.WriteString(Seperator)

	f.Close()
}

func (delegateTx DelegateTx) Prompt() {
	delegateOpt := delegateTx.DelegateOpt
	keyName := delegateTx.KeyName

	fmt.Print("\nDelegate Transaction\n")
	fmt.Print("\nKeyname: " + keyName + "\n")
	fmt.Print("\nDelegate Option\n\n")
	fmt.Printf("%+v\n", delegateOpt)
	fmt.Print(Seperator)
}
