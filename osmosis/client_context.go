package osmosis

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	osmosis "github.com/osmosis-labs/osmosis/app"
)

func GetDefaultClientContext() client.Context {
	// get osmosis codec
	encodingConfig := osmosis.MakeEncodingConfig()

	// set fields to default value
	defaultClientCtx := client.Context{
		JSONMarshaler:     encodingConfig.Marshaler,
		InterfaceRegistry: encodingConfig.InterfaceRegistry,
		TxConfig:          encodingConfig.TxConfig,
		LegacyAmino:       encodingConfig.Amino,
		Input:             os.Stdin,
		AccountRetriever:  authtypes.AccountRetriever{},
		BroadcastMode:     flags.BroadcastBlock,
		HomeDir:           HomeDir,
		OutputFormat:      "json",
		Simulate:          false,
		KeyringDir:        HomeDir,
		ChainID:           ChainId,
		NodeURI:           Node,
		GenerateOnly:      false,
		Offline:           false,
		UseLedger:         false,
		SkipConfirm:       true,
		SignModeStr:       "",
	}

	// keyring and client setting
	kr, _ := client.NewKeyringFromBackend(defaultClientCtx, "os")
	client, _ := client.NewClientFromNode(Node)

	defaultClientCtx = defaultClientCtx.WithKeyring(kr).
		WithClient(client).
		WithViper("OSMOSIS")
	return defaultClientCtx
}
