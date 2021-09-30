package cli

import (
	"github.com/notional-labs/cookiemonster/invest"
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
				err := investment.Invest()
				if err != nil {
					return err
				}
			}
			return nil
		},
	}

	return cmd
}
