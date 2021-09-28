// step 2 is pooling
package phase

import (
	"fmt"
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
	Name        string
	Config      map[string]int
	ConfigDenom string
}

// map from pool id to pooling amount in uosmo
type MapFromPoolToAmount map[int]*big.Int

// create txs from a strategy
func (poolStrategy PoolStrategy) GetTransactions(keyName string) (transaction.Transactions, error) {
	mapFromPoolToAmount, err := poolStrategy.GetMapFromPoolToAmount(keyName)
	if err != nil {
		return nil, err
	}
	BatchPool(mapFromPoolToAmount, keyName)

}

func (poolStrategy PoolStrategy) GetMapFromPoolToAmount(keyName string) (MapFromPoolToAmount, error) {
	uosmoBalance, err := query.QueryUosmoBalance(keyName)
	if err != nil {
		return nil, err
	}
	if poolStrategy.Name == "greedy" {
		fmt.Println("do something")
	}

	if poolStrategy.ConfigDenom == "percentages" {
		mapFromPoolToAmount := MapFromPoolToAmount{}
		for poolIdString, percentage := range poolStrategy.Config {
			poolId, err := strconv.Atoi(poolIdString)
			if err != nil {
				return nil, err
			}
			poolAmountUosmo := &big.Int{}
			poolAmountUosmo.Mul(big.NewInt(int64(percentage)), uosmoBalance)
			poolAmountUosmo.Div(poolAmountUosmo, big.NewInt(100))
			mapFromPoolToAmount[poolId] = poolAmountUosmo
		}
		return mapFromPoolToAmount, nil
	} else if poolStrategy.ConfigDenom == "osmo" {
		mapFromPoolToAmount := MapFromPoolToAmount{}
		for poolIdString, poolAmountOsmo := range poolStrategy.Config {
			poolId, err := strconv.Atoi(poolIdString)
			if err != nil {
				return nil, err
			}
			poolAmountUosmo := &big.Int{}
			temp := &big.Int{}
			poolAmountUosmo.Mul(big.NewInt(int64(poolAmountOsmo)), temp.SetUint64(10e17))
			mapFromPoolToAmount[poolId] = poolAmountUosmo
		}
		return mapFromPoolToAmount, nil
	} else {
		return nil, fmt.Errorf("unknown config denom")

	}
}

func BatchPool(mapFromPoolToAmount MapFromPoolToAmount, keyName string) ([]transaction.Transaction, error) {
	transactions := transaction.Transactions{}
	for poolId, poolAmount := range mapFromPoolToAmount {
		swapTx, err := SwapHalfAmountToPool(poolId, poolAmount, keyName)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, swapTx)

	}
}

// pool using the specified amount of osmo
func Pool(poolId int, uosmoAmount *big.Int, keyName string) {

}

// swap half the osmo amount to aonther token in the pool of specified id
func SwapHalfAmountToPool(poolId int, uosmoAmount *big.Int, keyName string) (transaction.Transaction, error) {
	halfUosmoAmount := &big.Int{}
	halfUosmoAmount.Div(uosmoAmount, big.NewInt(2))

	pool, err := query.QueryPoolId(poolId)
	if err != nil {
		return nil, err
	}

	tokenOutDenom := pool.PoolAssets[0].Token.Denom
	// uosmo price in tokenOut
	tokenOutPrice, err := query.QuerySpotPrice(poolId, tokenOutDenom, "uosmo")
	if err != nil {
		return nil, err
	}
	tokenOutAmount := BigIntMulFloat(halfUosmoAmount, tokenOutPrice)

	swapFeePercentage := pool.PoolParams.SwapFee
	swapFeeAmount := BigIntMulSDKDec(tokenOutAmount, swapFeePercentage)

	tokenOutAmount.Sub(tokenOutAmount, swapFeeAmount)

	slippageAmount := BigIntMulFloat(tokenOutAmount, slippagePercentage)

	tokenOutAmount.Sub(tokenOutAmount, slippageAmount)

	swapOpt := transaction.SwapOption{
		SwapRoutePoolIds: []int{poolId},
		SwapRouteDenoms:  []string{tokenOutDenom},
		TokenInAmount:    sdk.NewIntFromBigInt(halfUosmoAmount),
		TokenInDenom:     "uosmo",
		TokenOutMinAmt:   sdk.NewIntFromBigInt(tokenOutAmount),
	}
	swapTx := transaction.SwapTx{
		SwapOpt: swapOpt,
		KeyName: keyName,
	}
	return swapTx, nil
}

func BigIntMulFloat(x *big.Int, y float64) *big.Int {
	yDec, _ := sdk.NewDecFromStr(strconv.FormatFloat(y, 'f', -1, 64))
	return BigIntMulSDKDec(x, yDec)
}

func BigIntMulSDKDec(x *big.Int, y sdk.Dec) *big.Int {
	yBigInt := y.BigInt()

	temp := &big.Int{}
	z := &big.Int{}
	z.Mul(x, yBigInt)
	z.Div(x, temp.SetUint64(10e17))
	return z
}
