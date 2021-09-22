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

func NewMsgSwapExactAmountIn(swapOpt SwapOption, txOpt TxOption) (sdk.Msg, error) {
	routes, err := swapAmountInRoutes(swapOpt)
	if err != nil {
		return nil, err
	}

	tokenIn := sdk.Coin{Denom: swapOpt.tokenInDenom, Amount: swapOpt.tokenOutMinAmt}

	tokenOutMinAmt := swapOpt.tokenOutMinAmt

	msg := &types.MsgSwapExactAmountIn{
		Sender:            txOpt.FromAddr.String(),
		Routes:            routes,
		TokenIn:           tokenIn,
		TokenOutMinAmount: tokenOutMinAmt,
	}
	return msg, nil
}

func Swap(txOpt TxOption, swapOpt SwapOption) error {
	// build msg for tx
	msg, err := NewMsgSwapExactAmountIn(swapOpt, txOpt)
	if err != nil {
		return err
	}
	// build tx context
	clientCtx := client.Context{}
	SetContextFromTxOption(clientCtx, txOpt)
	txf := NewFactoryCLI(clientCtx)
	return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
}
