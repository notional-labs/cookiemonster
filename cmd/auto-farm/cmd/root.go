package cmd

import (
	"context"
	"os"
	"path/filepath"

	"github.com/osmosis-labs/osmosis/app/params"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	investcli "github.com/notional-labs/cookiemonster/invest/cli"
	osmosisapp "github.com/osmosis-labs/osmosis/app"
	"github.com/spf13/cobra"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	// "github.com/spf13/viper"
)

// NewRootCmd creates a new root command for simd. It is called once in the
// main function.
func NewRootCmd() (*cobra.Command, params.EncodingConfig) {
	encodingConfig := osmosisapp.MakeEncodingConfig()
	initClientCtx := client.Context{}.
		WithJSONMarshaler(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(types.AccountRetriever{}).
		WithBroadcastMode(flags.BroadcastBlock).
		WithHomeDir(DefaultNodeHome()).
		WithViper("OSMOSIS")

	rootCmd := &cobra.Command{
		Use:   "cookiemonster",
		Short: "tool for auto pooling and staking to osmosis",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			initClientCtx = client.ReadHomeFlag(initClientCtx, cmd)

			initClientCtx, err := config.ReadFromClientConfig(initClientCtx)
			if err != nil {
				return err
			}

			if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}

			return server.InterceptConfigsPreRunHandler(cmd)
		},
	}

	initRootCmd(rootCmd, encodingConfig)

	return rootCmd, encodingConfig
}

func initRootCmd(rootCmd *cobra.Command, encodingConfig params.EncodingConfig) {
	authclient.Codec = encodingConfig.Marshaler

	cfg := sdk.GetConfig()
	cfg.Seal()

	// add keybase
	rootCmd.AddCommand(
		investcli.NewInvestCmd(),
		keys.Commands(DefaultNodeHome()),
	)
}

// Execute executes the root command.
func Execute(rootCmd *cobra.Command) error {
	// Create and set a client.Context on the command's Context. During the pre-run
	// of the root command, a default initialized client.Context is provided to
	// seed child command execution with values such as AccountRetriver, Keyring,
	// and a Tendermint RPC. This requires the use of a pointer reference when
	// getting and setting the client.Context. Ideally, we utilize
	// https://github.com/spf13/cobra/pull/1118.
	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &client.Context{})
	ctx = context.WithValue(ctx, server.ServerContextKey, server.NewDefaultContext())

	executor := tmcli.PrepareBaseCmd(rootCmd, "", osmosisapp.DefaultNodeHome)
	return executor.ExecuteContext(ctx)
}

func DefaultNodeHome() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	DefaultNodeHome := filepath.Join(userHomeDir, ".auto-invest")
	return DefaultNodeHome
}
