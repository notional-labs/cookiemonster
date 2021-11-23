package cli

import (
	"fmt"
	"math/big"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/notional-labs/cookiemonster/api"
	cmdquery "github.com/notional-labs/cookiemonster/command/query"
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

func NewAutoInvestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auto-invest [path_to_investments_json]",
		Short: "pool and stake on osmosis using instruction from a json file, do so every epoch time",
		RunE: func(cmd *cobra.Command, args []string) error {
			api.InitAPI()
			for {

				pathToInvestmentJson := args[0]
				investments, err := invest.LoadInvestmentsFromFile(pathToInvestmentJson)
				if err != nil {
					return err
				}

				// report path is the path to tranasaction report file
				reportPath, _ := cmd.Flags().GetString(FlagReport)

				for _, investment := range investments {
					go func(investment *invest.Investment) error {
						keyName := investment.KeyName
						uosmoBalance, err := cmdquery.QueryUosmoBalance(cmd, keyName)
						if err != nil {
							return err
						}
						if uosmoBalance.Cmp(big.NewInt(1000000)) > 0 {
							err = investment.Invest(cmd, reportPath)
							if err != nil {
								return err
							}
						}
						return nil
					}(&investment)
				}
				fmt.Println(1)
				time.Sleep(5 * time.Minute)
			}
		},
	}
	cmd.Flags().String(FlagReport, "", "path to transaction report")
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
