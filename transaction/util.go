package transaction

import (
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)

// set fields in client context with keyname
func SetKeyNameToContext(clientCtx client.Context, keyName string) (client.Context, error) {
	from := keyName
	fromAddr, fromName, _, err := client.GetFromFields(clientCtx.Keyring, from, clientCtx.GenerateOnly)
	if err != nil {
		return clientCtx, err
	}
	clientCtx = clientCtx.WithFrom(from).WithFromAddress(fromAddr).WithFromName(fromName)

	return clientCtx, nil
}

func NewTxFactoryFromClientCtx(clientCtx client.Context) tx.Factory {
	transactionFactory := tx.Factory{}
	// set fields from clientCtx
	transactionFactory = transactionFactory.WithTxConfig(clientCtx.TxConfig).
		WithAccountRetriever(clientCtx.AccountRetriever).
		WithKeybase(clientCtx.Keyring).
		WithChainID(clientCtx.ChainID)

	// default value
	accNum := 0
	accSeq := 0
	gasAdj := 1
	memo := ""
	timeoutHeight := 0
	gasStr := ""
	gasSetting, _ := flags.ParseGasSetting(gasStr)
	feesStr := ""
	gasPricesStr := ""
	signMode := signing.SignMode_SIGN_MODE_DIRECT

	// set fields to default value
	transactionFactory = transactionFactory.WithGas(gasSetting.Gas).
		WithSimulateAndExecute(gasSetting.Simulate).
		WithAccountNumber(uint64(accNum)).
		WithSequence(uint64(accSeq)).
		WithTimeoutHeight(uint64(timeoutHeight)).
		WithGasAdjustment(float64(gasAdj)).
		WithMemo(memo).
		WithSignMode(signMode).
		WithFees(feesStr).
		WithGasPrices(gasPricesStr)

	return transactionFactory
}
