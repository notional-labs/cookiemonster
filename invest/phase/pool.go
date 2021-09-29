// step 2 is pooling
package phase

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
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
func (poolStrategy PoolStrategy) MakeTransactions(keyName string, totalPoolAmount *big.Int) transaction.Transactions {
	mapFromPoolToUosmoAmount, err := poolStrategy.MakeMapFromPoolToUosmoAmount(keyName, totalPoolAmount)
	if err != nil {
		return nil
	}
	transactionBatch := transaction.Transactions{}
	for poolId, poolUosmoAmount := range mapFromPoolToUosmoAmount {
		swapAndPoolTx := NewSwapAndPoolTransaction(poolId, sdk.NewIntFromBigInt(poolUosmoAmount), keyName)
		transactionBatch = append(transactionBatch, swapAndPoolTx)
	}
	return transactionBatch
}

func (poolStrategy PoolStrategy) MakeMapFromPoolToUosmoAmount(keyName string, totalPoolAmount *big.Int) (MapFromPoolToUosmoAmount, error) {
	if poolStrategy.Name == "greedy" {
		fmt.Println("do something")
	}

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

func loadPoolStrategy(fileLocation string) (*[]PoolStrategy, error) {
	file, err := os.Open(fileLocation)
	if (err!=nil) {
		fmt.Println("Unable to open json at "+fileLocation)
		return nil, err
	}
	reader := bufio.NewReader(file)
	jsonData, _ := ioutil.ReadAll(reader)

	var strategies []PoolStrategy;
	jsonErr := json.Unmarshal(jsonData, &strategies)
	if jsonErr != nil {
		fmt.Println("Unable to map JSON at "+fileLocation+" to a PoolStrategy");
		return nil, jsonErr
	}
	return &strategies, nil
}

func NewSwapAndPoolTransaction(poolId int, uosmoAmount sdk.Int, keyName string) transaction.Transaction {

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
