package cli

import (
	"fmt"
	"math/big"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/notional-labs/cookiemonster/invest"
	"github.com/notional-labs/cookiemonster/osmosis"
	"github.com/notional-labs/cookiemonster/query"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	FlagReport = "report"
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

			// report path is the path to tranasaction report file
			SetNode(cmd.Flags())
			for _, investment := range investments {
				err := investment.Invest()
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.Flags().String(flags.FlagNode, osmosis.Node, "<host>:<port> to Tendermint RPC interface for this chain")
	// cmd.Flags().String(FlagReport, "", "path to transaction report")
	return cmd
}

func NewAutoInvestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auto-invest [path_to_investments_json]",
		Short: "pool and stake on osmosis using instruction from a json file, do so every epoch time",
		RunE: func(cmd *cobra.Command, args []string) error {
			for {
				pathToInvestmentJson := args[0]
				investments, err := invest.LoadInvestmentsFromFile(pathToInvestmentJson)
				if err != nil {
					return err
				}

				// report path is the path to tranasaction report file
				SetNode(cmd.Flags())
				for id := range investments {
					go func(investment *invest.Investment) error {
						keyName := investment.KeyName
						fmt.Println(keyName)
						uosmoBalance, err := query.QueryUosmoBalance(keyName)
						if err != nil {
							return err
						}

						if uosmoBalance.Cmp(big.NewInt(1000000)) > 0 {
							err = investment.Invest()
							if err != nil {
								return err
							}
						}
						return nil
					}(&investments[id]) //nolint:errcheck
				}
				time.Sleep(5 * time.Minute)
			}
		},
	}
	// cmd.Flags().String(FlagReport, "", "path to transaction report")
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd

}

func SetNode(flagSet *pflag.FlagSet) {
	if flagSet.Changed(flags.FlagNode) {
		osmosis.Node, _ = flagSet.GetString(flags.FlagNode)
	}
}
