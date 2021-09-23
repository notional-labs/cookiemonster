// step 2 is pooling
package phase

import (
	"github.com/notional-labs/cookiemonster/internal/query"

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
	O





}