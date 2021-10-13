package invest

import (
	"github.com/spf13/cobra"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/cookiemonster/query"
	"github.com/notional-labs/cookiemonster/transaction"
)

func MakeLockTxs(cmd *cobra.Command, keyName string, duration string) (transaction.Txs, error) {
	balances, err := query.QueryBalances(cmd, keyName)
	if err != nil {
		return nil, err
	}

	gammCoins := sdk.Coins{}
	for _, coin := range balances {
		if coin.Denom[:4] == "gamm" && coin.Amount.BigInt().Cmp(big.NewInt(0)) != 0 {
			gammCoins = append(gammCoins, coin)
		}
	}

	lockTxs := transaction.Txs{}
	for _, gammCoin := range gammCoins {
		lockOpt := transaction.LockOption{
			Denom:    gammCoin.Denom,
			Amount:   gammCoin.Amount,
			Duration: duration,
		}
		lockTx := transaction.LockTx{
			KeyName: keyName,
			LockOpt: lockOpt,
		}

		lockTxs = append(lockTxs, lockTx)
	}

	return lockTxs, nil
}
