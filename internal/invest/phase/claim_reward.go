// step 1 is claim reward from staking
package phase

import (
	"github.com/notional-labs/cookiemonster/internal/invest"
	"github.com/notional-labs/cookiemonster/internal/transaction"
)

func ClaimRewardFromFund(fund invest.Fund) error {

	txOpt := transaction.TxOption{
		KeyName: fund.KeyName,
	}

	transaction.ClaimReward()

}
