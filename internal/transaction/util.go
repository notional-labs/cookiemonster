package transaction

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"

	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)

func SetContextFromTxOption(clientCtx client.Context, txOpt TxOption) (client.Context, error) {
	clientCtx, err := SetBasicContextFromTxOption(clientCtx, txOpt)
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
		from := txOpt.KeyName
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

func SetBasicContextFromTxOption(clientCtx client.Context, txOpt TxOption) (client.Context, error) {
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
		chainID := txOpt.ChainId
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
		rpcURI := txOpt.Node
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

	f := tx.Factory{}

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

	f.WithTxConfig(clientCtx.TxConfig)
	f.WithAccountRetriever(clientCtx.AccountRetriever)
	f.WithKeybase(clientCtx.Keyring)
	f.WithChainID(clientCtx.ChainID)
	f.WithGas(gasSetting.Gas)
	f.WithSimulateAndExecute(gasSetting.Simulate)
	f.WithAccountNumber(uint64(accNum))
	f.WithSequence(uint64(accSeq))
	f.WithTimeoutHeight(uint64(timeoutHeight))
	f.WithGasAdjustment(float64(gasAdj))
	f.WithMemo(memo)
	f.WithSignMode(signMode)

	feesStr := ""
	f = f.WithFees(feesStr)

	gasPricesStr := ""
	f = f.WithGasPrices(gasPricesStr)

	return f
}
