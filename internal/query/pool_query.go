package query

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	gammcli "github.com/osmosis-labs/osmosis/x/gamm/client/cli"
	"github.com/osmosis-labs/osmosis/x/gamm/types"
)

func QuerySpotPrice(poolId int, tokenInDenom string, tokenOutDenom string) (float64, error) {
	cmd := gammcli.GetCmdSpotPrice()
	clientCtx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return 0, err
	}
	queryClient := types.NewQueryClient(clientCtx)

	if err != nil {
		return 0, err
	}

	res, err := queryClient.SpotPrice(cmd.Context(), &types.QuerySpotPriceRequest{
		PoolId:        uint64(poolId),
		TokenInDenom:  tokenInDenom,
		TokenOutDenom: tokenOutDenom,
	})
	if err != nil {
		return 0, err
	}

	spotPriceString := res.GetSpotPrice()
	spotPrice, _ := strconv.ParseFloat(spotPriceString, 64)
	return spotPrice, nil
}
