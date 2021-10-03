package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/notional-labs/cookiemonster/invest"
	"github.com/notional-labs/cookiemonster/osmosis"
	"github.com/spf13/cobra"
)

func NewInvestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invest [path_to_investments_json]",
		Short: "pool and stake on osmosis using instruction from a json file",
		RunE: func(cmd *cobra.Command, args []string) error {
			pathToInvestmentJson := args[0]
			investments, err := invest.LoadInvestmentsFromFile(pathToInvestmentJson)
			if err != nil {
				return err
			}
			for _, investment := range investments {
				fmt.Printf("%+v\n", investment)
				err := investment.Invest()
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.Flags().String(flags.FlagNode, osmosis.Node, "<host>:<port> to Tendermint RPC interface for this chain")

	return cmd
}
