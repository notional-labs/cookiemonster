// step 2 is pooling
package phase

import (
	"math/big"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/cookiemonster/query"
	"github.com/notional-labs/cookiemonster/transaction"
)

var (
	slippagePercentage float64 = 0.003
)

type PoolStrategy struct {
	Name       string
	Config     map[string]int
	ConfigType string
}

// map from pool id to pooling amount
type PoolAmountMap map[int]*big.Int

// create txs from a strategy
func (poolStrategy PoolStrategy) CreateTxsFromStrategy(keyName string) ([]transaction.Transaction, error) {

	osmoBalance, err := query.QueryOsmoBalance(keyName)
	if err != nil {
		return nil, err
	}

	poolAmountsMap := PoolAmountMap{}
	if poolStrategy.ConfigType == "percentages" {
		for poolIdString, percentage := range poolStrategy.Config {
			poolId, err := strconv.Atoi(poolIdString)
			if err != nil {
				return nil, err 
			}
			poolAmount := &big.Int{}
			poolAmount.Mul(big.NewInt(int64(percentage)), osmoBalance)
			poolAmount.Div(poolAmount, big.NewInt(100))
			poolAmountsMap[poolId] = poolAmount
		}
	}
	BatchPool(poolAmountsMap, keyName)

}

func CreatePoolAmountMapFromStrategy

func BatchPool(poolAmountsMap PoolAmountMap, keyName string) error{
	for poolId, poolAmount := range poolAmountsMap {
		err := SwapHalfAmountToPool(poolId, poolAmount, keyName)
		if err != nil {
			return err
		}
	}
}

// pool using the specified amount of osmo
func Pool(poolId int, osmoAmount *big.Int, keyName string) {

}

// swap half the osmo amount to aonther token in the pool of specified id
func SwapHalfAmountToPool(poolId int, osmoAmount *big.Int, keyName string) error {
	halfOsmoAmount := osmoAmount.Div(osmoAmount, big.NewInt(2))

	pool, err := query.QueryPoolId(poolId)
	if err != nil {
		return err
	}

	tokenOutDenom := pool.PoolAssets[0].Token.Denom
	// uosmo price in tokenOut
	tokenOutPrice, err := query.QuerySpotPrice(poolId, tokenOutDenom, "uosmo")
	if err != nil {
		return err
	}
	tokenOutAmount := BigIntMulFloat(halfOsmoAmount, tokenOutPrice)

	swapFeePercentage := pool.PoolParams.SwapFee
	swapFeeAmount := BigIntMulSDKDec(tokenOutAmount, swapFeePercentage)

	tokenOutAmount.Sub(tokenOutAmount, swapFeeAmount)

	slippageAmount := BigIntMulFloat(tokenOutAmount, slippagePercentage)

	tokenOutAmount.Sub(tokenOutAmount, slippageAmount)

	swapOpt := transaction.SwapOption{
		SwapRoutePoolIds: []int{poolId},
		SwapRouteDenoms:  []string{tokenOutDenom},
		TokenInAmount:    sdk.NewIntFromBigInt(halfOsmoAmount),
		TokenInDenom:     "uosmo",
		TokenOutMinAmt:   sdk.NewIntFromBigInt(tokenOutAmount),
	}
	err = transaction.Swap(keyName, swapOpt)
	if err != nil {
		return err
	}
	return nil
}

func BigIntMulFloat(x *big.Int, y float64) *big.Int {
	yDec, _ := sdk.NewDecFromStr(strconv.FormatFloat(y, 'f', -1, 64))
	yBigInt := yDec.BigInt()

	temp := &big.Int{}
	z := &big.Int{}
	z.Mul(x, yBigInt)
	z.Div(x, temp.SetUint64(10e17))
	return z
}

func BigIntMulSDKDec(x *big.Int, y sdk.Dec) *big.Int {
	yBigInt := y.BigInt()

	temp := &big.Int{}
	z := &big.Int{}
	z.Mul(x, yBigInt)
	z.Div(x, temp.SetUint64(10e17))
	return z
}
