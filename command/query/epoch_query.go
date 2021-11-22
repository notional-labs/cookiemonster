package query

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/osmosis-labs/osmosis/x/epochs/types"
	"github.com/spf13/cobra"
)

func QueryEpoch(cmd *cobra.Command) (*types.QueryEpochsInfoResponse, error) {
	clientCtx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return nil, err
	}
	queryClient := types.NewQueryClient(clientCtx)

	res, err := queryClient.EpochInfos(cmd.Context(), &types.QueryEpochsInfoRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}
