package main

import (
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"fmt"

	osmosis "github.com/osmosis-labs/osmosis/app"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"

	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)

type BankSendOption struct {
	ToAddr sdk.AccAddress
	Denom  string
	Amount sdk.Int
}

func BankSend(keyName string, sendOpt BankSendOption) error {
	// build tx context
	encodingConfig := osmosis.MakeEncodingConfig()
	initClientCtx := client.Context{}.
		WithJSONMarshaler(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithBroadcastMode(flags.BroadcastBlock).
		WithHomeDir(osmosis.DefaultNodeHome).
		WithViper("OSMOSIS")

	clientCtx, err := SetContextFromKeyName(initClientCtx, keyName)
	fmt.Println(clientCtx.KeyringDir)
	fmt.Println(clientCtx.Keyring)
	if err != nil {
		return err
	}
	txf := NewFactoryCLI(initClientCtx)

	// build msg for tx
	toAddr := sendOpt.ToAddr
	fromAddr := clientCtx.GetFromAddress()
	coin := sdk.Coin{Denom: sendOpt.Denom, Amount: sendOpt.Amount}
	coins := sdk.Coins([]sdk.Coin{coin})
	msg := types.NewMsgSend(fromAddr, toAddr, coins)

	return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
}

func main() {
	toAddr, _ := sdk.AccAddressFromBech32("osmo13k9w2pexxtyyfuw7fmxh6rpwwl9udxkk26nfle")
	bankSendOpt := BankSendOption{
		ToAddr: toAddr,
		Denom:  "stake",
		Amount: sdk.NewInt(12),
	}
	fmt.Println(BankSend("april", bankSendOpt))

}

func SetContextFromKeyName(clientCtx client.Context, keyName string) (client.Context, error) {
	clientCtx, err := SetBasicContextFromKeyName(clientCtx, keyName)
	if err != nil {
		return clientCtx, err
	}
	if !clientCtx.GenerateOnly {
		genOnly := false
		clientCtx = clientCtx.WithGenerateOnly(genOnly)
	}
	if !clientCtx.Offline {
		offline := false
		clientCtx = clientCtx.WithOffline(offline)
	}
	if !clientCtx.UseLedger {
		useLedger := false
		clientCtx = clientCtx.WithUseLedger(useLedger)
	}
	if clientCtx.BroadcastMode == "" {
		bMode := "sync"
		clientCtx = clientCtx.WithBroadcastMode(bMode)
	}
	if !clientCtx.SkipConfirm {
		skipConfirm := true
		clientCtx = clientCtx.WithSkipConfirmation(skipConfirm)
	}
	if clientCtx.SignModeStr == "" {
		signModeStr := ""
		clientCtx = clientCtx.WithSignModeStr(signModeStr)
	}

	// if clientCtx.FeeGranter == nil {
	// 	granter := ""

	// 	if granter != "" {
	// 		granterAcc, err := sdk.AccAddressFromBech32(granter)
	// 		if err != nil {
	// 			return clientCtx, err
	// 		}

	// 		clientCtx = clientCtx.WithFeeGranterAddress(granterAcc)
	// 	}
	// }

	if clientCtx.From == "" {
		from := keyName
		fromAddr, fromName, keyType, err := client.GetFromFields(clientCtx.Keyring, from, clientCtx.GenerateOnly)
		if err != nil {
			return clientCtx, err
		}
		clientCtx = clientCtx.WithFrom(from).WithFromAddress(fromAddr).WithFromName(fromName)
		// If the `from` signer account is a ledger key, we need to use
		// SIGN_MODE_AMINO_JSON, because ledger doesn't support proto yet.
		// ref: https://github.com/cosmos/cosmos-sdk/issues/8109
		if keyType == keyring.TypeLedger && clientCtx.SignModeStr != flags.SignModeLegacyAminoJSON {
			fmt.Println("Default sign-mode 'direct' not supported by Ledger, using sign-mode 'amino-json'.")
			clientCtx = clientCtx.WithSignModeStr(flags.SignModeLegacyAminoJSON)
		}
	}
	return clientCtx, nil
}

func SetBasicContextFromKeyName(clientCtx client.Context, keyName string) (client.Context, error) {
	if clientCtx.OutputFormat == "" {
		output := "json"
		clientCtx = clientCtx.WithOutputFormat(output)
	}
	if !clientCtx.Simulate {
		dryRun := false
		clientCtx = clientCtx.WithSimulation(dryRun)
	}
	if clientCtx.KeyringDir == "" {
		keyringDir := ""

		// The keyring directory is optional and falls back to the home directory
		// if omitted.
		if keyringDir == "" {
			keyringDir = clientCtx.HomeDir
		}

		clientCtx = clientCtx.WithKeyringDir(keyringDir)
	}

	if clientCtx.ChainID == "" {
		chainID := "osmosis-1"
		clientCtx = clientCtx.WithChainID(chainID)
	}
	if clientCtx.Keyring == nil {
		keyringBackend := "os"

		if keyringBackend != "" {
			kr, err := client.NewKeyringFromBackend(clientCtx, keyringBackend)
			if err != nil {
				return clientCtx, err
			}

			clientCtx = clientCtx.WithKeyring(kr)
		}
	}
	if clientCtx.Client == nil {
		rpcURI := "http://0.0.0.0:26657"
		if rpcURI != "" {
			clientCtx = clientCtx.WithNodeURI(rpcURI)

			client, err := client.NewClientFromNode(rpcURI)
			if err != nil {
				return clientCtx, err
			}

			clientCtx = clientCtx.WithClient(client)
		}
	}
	return clientCtx, nil
}

// NewFactoryCLI creates a new Factory.
func NewFactoryCLI(clientCtx client.Context) tx.Factory {
	signModeStr := clientCtx.SignModeStr

	transactionFactory := tx.Factory{}

	signMode := signing.SignMode_SIGN_MODE_UNSPECIFIED
	switch signModeStr {
	case flags.SignModeDirect:
		signMode = signing.SignMode_SIGN_MODE_DIRECT
	case flags.SignModeLegacyAminoJSON:
		signMode = signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON
	}

	// offline mode only
	accNum := 0
	accSeq := 0

	gasAdj := 1
	memo := ""
	timeoutHeight := 0
	gasStr := ""
	gasSetting, _ := flags.ParseGasSetting(gasStr)

	transactionFactory.WithTxConfig(clientCtx.TxConfig)
	transactionFactory.WithAccountRetriever(clientCtx.AccountRetriever)
	transactionFactory.WithKeybase(clientCtx.Keyring)
	transactionFactory.WithChainID(clientCtx.ChainID)
	transactionFactory.WithGas(gasSetting.Gas)
	transactionFactory.WithSimulateAndExecute(gasSetting.Simulate)
	transactionFactory.WithAccountNumber(uint64(accNum))
	transactionFactory.WithSequence(uint64(accSeq))
	transactionFactory.WithTimeoutHeight(uint64(timeoutHeight))
	transactionFactory.WithGasAdjustment(float64(gasAdj))
	transactionFactory.WithMemo(memo)
	transactionFactory.WithSignMode(signMode)

	feesStr := ""
	transactionFactory = transactionFactory.WithFees(feesStr)

	gasPricesStr := ""
	transactionFactory = transactionFactory.WithGasPrices(gasPricesStr)

	return transactionFactory
}
