package query

import (
	"context"

	"github.com/notional-labs/cookiemonster/osmosis"
	epoch "github.com/osmosis-labs/osmosis/x/epochs/types"
)

func QueryEpoch() (int64, error) {
	clientCtx := osmosis.GetDefaultClientContext()

	queryClient := epoch.NewQueryClient(clientCtx)

	res, err := queryClient.CurrentEpoch(context.Background(), &epoch.QueryCurrentEpochRequest{
		Identifier: "day",
	})
	if err != nil {
		return -1, err
	}
	currentEpoch := res.CurrentEpoch
	return currentEpoch, nil
}

// func QueryEpochTime() {
// 	clientCtx := osmosis.GetDefaultClientContext()

// 	queryClient := epoch.NewQueryClient(clientCtx)

// 	res, err := queryClient.EpochInfos(context.Background(), &epoch.QueryEpochsInfoRequest{})

// 	if err != nil {
// 		return -1, err
// 	}

// 	epochInfos := res.GetEpochs()

// 	for epochInfo := ra

// }
