package cli

import (
	"math/big"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"golang.org/x/sync/errgroup"

	"github.com/notional-labs/cookiemonster/invest"
	"github.com/notional-labs/cookiemonster/command/query"
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
				reportPath, _ := cmd.Flags().GetString(FlagReport)
				eg := new(errgroup.Group)

				eg.Go(func() error {
					for _, investment := range investments {
						if err := retry.Do(func() error {
							keyName := investment.KeyName
							uosmoBalance, err := query.QueryUosmoBalance(cmd, keyName)
							if err != nil {
								return err
							}
							if uosmoBalance.Cmp(big.NewInt(1000000)) == 1 {
								err := investment.Invest(cmd, reportPath)
								if err != nil {
									return err
								}
							return nil
							}}, retry.Attempts(5), retry.Delay(time.Millisecond*500), retry.LastErrorOnly(true)); err != nil {
								return err
							}
							time.Sle
						})
					}
				})
				

				

				}
			}
		},
	}
	cmd.Flags().String(FlagReport, "", "path to transaction report")
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}