// step 2 is pooling
package phase

import (
	"github.com/notional-labs/cookiemonster/query"

)

type PoolStrategy struct {
	Name   string
	Config map[string]string
}

// create txs from a strategy
func(poolStrategy PoolStrategy) CreateTxs(keyName string) error {
	OsmoBalance, err := query.QueryOsmoBalance()
	if err != nil {
		return err
	}

	OsmoAmount := OsmoBalance.Amount.BigInt()



	for _, coin range balances {
		if coin 

	}
}


func(poolStrategy PoolStrategy) {
	
}

// package mapping

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"os"
// )

// var (
// 	MapFromDenomToPoolId map[string]string
// )

// func GetMapFromDenomToPoolId() error {

// 	if err != nil {
// 		return err
// 	}
// 	// if we os.Open returns an error then handle it

// 	// defer the closing of our jsonFile so that we can parse it later on
// 	byteValue, err := ioutil.ReadAll(jsonFile)
// 	if err != nil {
// 		return err
// 	}
// 	var result map[string]string
// 	json.Unmarshal([]byte(byteValue), &result)

// 	ChainId = result["ChainId"]
// 	Node = result["Node"]
// 	return nil
// }
