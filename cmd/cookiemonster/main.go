package main

import (
	"os"

	"github.com/notional-labs/cookiemonster/cmd/cookiemonster/cmd"
	"github.com/notional-labs/cookiemonster/frontend"
	"github.com/osmosis-labs/osmosis/app/params"
	// "github.com/osmosis-labs/osmosis/app/params"
)

// This will start cookie monster and serve the front end.  It also might not be quite exactly designed right.
// Let's consider dropping cobra, it makes things complicated and there are better equivalent libs.
func main() {
	params.SetAddressPrefixes()
	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	go frontend.Serve()
}
