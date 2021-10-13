package cli

import (
	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/notional-labs/cookiemonster/invest"
	"github.com/spf13/cobra"
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

			for _, investment := range investments {
				err := investment.Invest(cmd, reportPath)
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
