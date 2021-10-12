package main

import (
	"os"

	"github.com/notional-labs/cookiemonster/cmd/auto-farm/cmd"
	"github.com/osmosis-labs/osmosis/app/params"
	// "github.com/osmosis-labs/osmosis/app/params"
)

func main() {
	params.SetAddressPrefixes()
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
