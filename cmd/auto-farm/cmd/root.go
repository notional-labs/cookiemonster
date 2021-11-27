package cmd

import (

	// "github.com/osmosis-labs/osmosis/app/params"

	"github.com/cosmos/cosmos-sdk/client/keys"
	apicli "github.com/notional-labs/cookiemonster/api/cli"
	investcli "github.com/notional-labs/cookiemonster/invest/cli"
	"github.com/notional-labs/cookiemonster/osmosis"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "auto-invest",
		Short: "auto-invest tool for osmosis",
	}
	rootCmd.AddCommand(
		apicli.NewInitApi(),
		investcli.NewInvestCmd(),
		investcli.NewAutoInvestCmd(),
		keys.Commands(osmosis.HomeDir),
	)

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
// func Execute(rootCmd *cobra.Command) {
// 	rootCmd.Execute()
// }

// func initRootCmd() {
// 	cobra.OnInitialize(initConfig)

// 	// Here you will define your flags and configuration settings.
// 	// Cobra supports persistent flags, which, if defined here,
// 	// will be global for your application.

// 	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.getrewards.yaml)")

// 	// Cobra also supports local flags, which will only run
// 	// when this action is called directly.
// 	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }

// // initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := os.UserHomeDir()
// 		cobra.CheckErr(err)

// 		// Search config in home directory with name ".getrewards" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigType("yaml")
// 		viper.SetConfigName(".getrewards")
// 	}

// 	viper.AutomaticEnv() // read in environment variables that match

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
// 	}
// }
