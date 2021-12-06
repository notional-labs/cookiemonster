package cli

import (
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/notional-labs/cookiemonster/invest"
	"github.com/notional-labs/cookiemonster/osmosis"
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
			reportPath, _ := cmd.Flags().GetString(FlagReport)
			SetNode(cmd.Flags())
			for _, investment := range investments {
				err := investment.Invest(reportPath)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.Flags().String(flags.FlagNode, osmosis.Node, "<host>:<port> to Tendermint RPC interface for this chain")
	cmd.Flags().String(FlagReport, "", "path to transaction report")
	return cmd
}
func NewInvestToDieCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invest_non_stop [path_to_investments_json]",
		Short: "Invest and non-stop",
		RunE: func(cmd *cobra.Command, args []string) error {

			pathToInvestmentJson := args[0]
			investments, err := invest.LoadInvestmentsFromFile(pathToInvestmentJson)
			if err != nil {
				return err
			}

			// report path is the path to tranasaction report file
			reportPath, _ := cmd.Flags().GetString(FlagReport)

			for _, investment := range investments {
				err := investment.InvestToDie(reportPath)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.Flags().String(FlagReport, "", "path to transaction report")
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func SetNode(flagSet *pflag.FlagSet) {
	if flagSet.Changed(flags.FlagNode) {
		osmosis.Node, _ = flagSet.GetString(flags.FlagNode)
	}
}
