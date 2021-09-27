package transaction

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/notional-labs/cookiemonster/osmosis"
	"gopkg.in/yaml.v3"
)

type DelegateOption struct {
	ValAddr sdk.ValAddress
	Denom   string
	Amount  sdk.Int
}

func Delegate(keyName string, delegateOpt DelegateOption) error {
	// build tx context
	clientCtx := osmosis.DefaultClientCtx
	clientCtx, err := SetKeyNameToContext(clientCtx, keyName)
	if err != nil {
		return err
	}

	txf := NewTxFactoryFromClientCtx(clientCtx)

	// build msg for tx
	valAddr := delegateOpt.ValAddr
	delAddr := clientCtx.FromAddress
	amount := sdk.Coin{Denom: delegateOpt.Denom, Amount: delegateOpt.Amount}
	msg := types.NewMsgDelegate(delAddr, valAddr, amount)

	return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
}

type UndelegateOption struct {
	ValAddr sdk.ValAddress
	Denom   string
	Amount  sdk.Int
}

func Undelegate(keyName string, undelegateOpt UndelegateOption) error {
	// build tx context
	clientCtx := osmosis.DefaultClientCtx
	clientCtx, err := SetKeyNameToContext(clientCtx, keyName)
	if err != nil {
		return err
	}

	txf := NewTxFactoryFromClientCtx(clientCtx)

	// build msg for tx
	valAddr := undelegateOpt.ValAddr
	delAddr := clientCtx.FromAddress
	amount := sdk.Coin{Denom: undelegateOpt.Denom, Amount: undelegateOpt.Amount}
	msg := types.NewMsgDelegate(delAddr, valAddr, amount)

	return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
}

type DelegateTx struct {
	KeyName     string
	DelegateOpt DelegateOption
}

func (delegateTx DelegateTx) Execute() error {

	keyName := delegateTx.KeyName
	delegateOpt := delegateTx.DelegateOpt
	err := Delegate(keyName, delegateOpt)
	return err
}

func (delegateTx DelegateTx) Report() {

	delegateOpt := delegateTx.DelegateOpt
	keyName := delegateTx.KeyName

	f, _ := os.OpenFile("report", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	f.WriteString("\nDelegate Transaction\n")
	f.WriteString("\nKeyname: " + keyName + "\n")
	f.WriteString("\nDelegate Option\n\n")

	txData, _ := yaml.Marshal(delegateOpt)
	_, _ = f.Write(txData)
	f.WriteString(transactionSeperator)

	f.Close()
}
