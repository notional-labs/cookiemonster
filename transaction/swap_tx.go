package transaction

import (
	"errors"
	"os"

	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/cookiemonster/osmosis"
	"github.com/osmosis-labs/osmosis/x/gamm/types"
	"gopkg.in/yaml.v3"
)

type SwapOption struct {
	SwapRoutePoolIds []int
	SwapRouteDenoms  []string
	TokenInAmount    sdk.Int
	TokenInDenom     string
	TokenOutMinAmt   sdk.Int
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

	tokenIn := sdk.Coin{Denom: swapOpt.TokenInDenom, Amount: swapOpt.TokenInAmount}

	tokenOutMinAmt := swapOpt.TokenOutMinAmt

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
	clientCtx := osmosis.DefaultClientCtx
	clientCtx, err := SetKeyNameToContext(clientCtx, keyName)
	if err != nil {
		return err
	}

	txf := NewTxFactoryFromClientCtx(clientCtx)

	// build msg for tx
	fromAddr := clientCtx.GetFromAddress()
	msg, err := NewMsgSwapExactAmountIn(fromAddr, swapOpt)
	if err != nil {
		return err
	}

	return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
}

type SwapTx struct {
	KeyName string
	SwapOpt SwapOption
}

func (swapTx SwapTx) Execute() error {

	keyName := swapTx.KeyName
	swapOpt := swapTx.SwapOpt
	err := Swap(keyName, swapOpt)
	return err
}

func (swapTx SwapTx) Report() {

	swapOpt := swapTx.SwapOpt
	keyName := swapTx.KeyName

	f, _ := os.OpenFile("report", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	f.WriteString("\nSwap Transaction\n")
	f.WriteString("\nKeyname: " + keyName + "\n")
	f.WriteString("\nSwap Option\n\n")

	txData, _ := yaml.Marshal(swapOpt)
	_, _ = f.Write(txData)
	f.WriteString(transactionSeperator)

	f.Close()
}
