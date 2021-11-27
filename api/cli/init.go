package cli

import (
	"github.com/notional-labs/cookiemonster/api"
	"github.com/spf13/cobra"
)

const (
	FlagReport = "report"
)

func NewInitApi() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "initapi",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {

			api.InitAPI()
			return nil
		},
	}
	return cmd
}
