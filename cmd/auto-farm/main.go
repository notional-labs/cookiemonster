package main

import (
	"os"

	"github.com/notional-labs/cookiemonster/cmd/auto-farm/cmd"
	"github.com/osmosis-labs/osmosis/app/params"
)

func main() {
	params.SetAddressPrefixes()
	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
