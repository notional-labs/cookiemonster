package transaction

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/osmosis-labs/osmosis/x/gamm/types"
)

type SwapOption struct {
	SwapRoutePoolIds []int
	SwapRouteDenoms  []string
	Route            []types.SwapAmountInRoute
	tokenInAmount    sdk.Int
	tokenInDenom     string
	tokenOutMinAmt   sdk.Int
}

func swapAmountInRoutes(swapOpt SwapOption) ([]types.SwapAmountInRoute, error) {
	swapRoutePoolIds := swapOpt.SwapRoutePoolIds
	swapRouteDenoms := swapOpt.SwapRouteDenoms

	if len(swapRoutePoolIds) != len(swapRouteDenoms) {
		return nil, errors.New("swap route pool ids and denoms mismatch")
	}
	routes := []types.SwapAmountInRoute{}
	for index, poolID := range swapRoutePoolIds {
		routes = append(routes, types.SwapAmountInRoute{
			PoolId:        uint64(poolID),
			TokenOutDenom: swapRouteDenoms[index],
		})
	}
	return routes, nil
}

func NewMsgSwapExactAmountIn(fromAddr sdk.AccAddress, swapOpt SwapOption) (sdk.Msg, error) {
	routes, err := swapAmountInRoutes(swapOpt)
	if err != nil {
		return nil, err
	}

	tokenIn := sdk.Coin{Denom: swapOpt.tokenInDenom, Amount: swapOpt.tokenOutMinAmt}

	tokenOutMinAmt := swapOpt.tokenOutMinAmt

	msg := &types.MsgSwapExactAmountIn{
		Sender:            fromAddr.String(),
		Routes:            routes,
		TokenIn:           tokenIn,
		TokenOutMinAmount: tokenOutMinAmt,
	}
	return msg, nil
}

func Swap(keyName string, swapOpt SwapOption) error {
	// build tx context
	clientCtx := client.Context{}
	SetContextFromKeyName(clientCtx, keyName)
	txf := NewFactoryCLI(clientCtx)

	// build msg for tx
	fromAddr := clientCtx.GetFromAddress()
	msg, err := NewMsgSwapExactAmountIn(fromAddr, swapOpt)
	if err != nil {
		return err
	}

	return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
}
