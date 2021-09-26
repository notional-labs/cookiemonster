package mapping

import (
	"math/big"

	query "github.com/notional-labs/cookiemonster/query"
)

var (
	MapFromDenomToPoolId      map[string]string
	GetMapFromDenomToIBCDenom map[string]string
)

func GetMapFromDenomsToPoolId() map[string]int {
	var mapFromDenomToPoolId map[string]int
	var mapFromPoolIdToAmount map[int]*big.Int
	pools, err := query.QueryPools()
	if err != nil {
		return nil
	}

	for _, pool := range pools {
		poolId := int(pool.Id)
		if err != nil {
			return nil
		}

		denom := ""
		amount := big.NewInt(0)
		for _, coin := range pool.PoolAssets {
			denom += coin.Token.Denom
			amount.Add(amount, coin.Token.Amount.BigInt())
		}
		mapFromPoolIdToAmount[poolId] = amount
		if idOfPoolWithSameDenom, ok := mapFromDenomToPoolId[denom]; ok {
			if mapFromPoolIdToAmount[poolId].Cmp(mapFromPoolIdToAmount[idOfPoolWithSameDenom]) == 1 {
				mapFromDenomToPoolId[denom] = poolId
			}
		} else {
			mapFromDenomToPoolId[denom] = poolId
		}
	}

	return mapFromDenomToPoolId
}

func GetMapFromDenomToIBCDenom() error {

}
