// step 2 is pooling
package invest

import (
	"fmt"
	"math/big"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/cookiemonster/transaction"
)

type PoolStrategy struct {
	Name        string
	Config      map[string]int
	ConfigDenom string
}

// map from pool id to uosmo amount to be pooled
type MapFromPoolToUosmoAmount map[int]*big.Int

// create txs from a strategy
func MakeSwapAndPoolTxs(keyName string, totalPoolAmount *big.Int, poolStrategy PoolStrategy) transaction.Txs {
	mapFromPoolToUosmoAmount, err := MakeMapFromPoolToUosmoAmount(totalPoolAmount, poolStrategy)

	if err != nil {
		return nil
	}
	txBatch := transaction.Txs{}
	for poolId, poolUosmoAmount := range mapFromPoolToUosmoAmount {
		swapAndPoolTx := NewSwapAndPooltx(poolId, sdk.NewIntFromBigInt(poolUosmoAmount), keyName)

		txBatch = append(txBatch, swapAndPoolTx)
	}
	return txBatch
}

// Make Map to decide which pool to pool and how much to pool in each pool
func MakeMapFromPoolToUosmoAmount(totalPoolAmount *big.Int, poolStrategy PoolStrategy) (MapFromPoolToUosmoAmount, error) {
	if poolStrategy.ConfigDenom == "percentages" {
		mapFromPoolToUosmoAmount := MapFromPoolToUosmoAmount{}
		for poolIdString, percentage := range poolStrategy.Config {
			poolId, err := strconv.Atoi(poolIdString)
			if err != nil {
				return nil, err
			}
			poolAmountUosmo := &big.Int{}
			poolAmountUosmo.Mul(big.NewInt(int64(percentage)), totalPoolAmount)
			poolAmountUosmo.Div(poolAmountUosmo, big.NewInt(100))
			mapFromPoolToUosmoAmount[poolId] = poolAmountUosmo
		}
		return mapFromPoolToUosmoAmount, nil
	} else if poolStrategy.ConfigDenom == "osmo" {
		mapFromPoolToUosmoAmount := MapFromPoolToUosmoAmount{}
		for poolIdString, poolAmountOsmo := range poolStrategy.Config {
			poolId, err := strconv.Atoi(poolIdString)
			if err != nil {
				return nil, err
			}
			poolAmountUosmo := &big.Int{}
			temp := &big.Int{}
			poolAmountUosmo.Mul(big.NewInt(int64(poolAmountOsmo)), temp.SetUint64(1e18))
			mapFromPoolToUosmoAmount[poolId] = poolAmountUosmo
		}
		return mapFromPoolToUosmoAmount, nil
	} else {
		return nil, fmt.Errorf("unknown config denom")
	}
}

func NewSwapAndPooltx(poolId int, uosmoAmount sdk.Int, keyName string) transaction.Tx {

	swapAndPoolOpt := transaction.SwapAndPoolOption{
		TokenInAmount:     uosmoAmount,
		PoolId:            uint64(poolId),
		TokenInDenom:      "uosmo",
		ShareOutMinAmount: sdk.NewInt(1),
	}

	swapAndPoolTx := transaction.SwapAndPoolTx{
		KeyName:        keyName,
		SwapAndPoolOpt: swapAndPoolOpt,
	}

	return swapAndPoolTx
}
