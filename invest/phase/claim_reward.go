// step 1 is claim reward from staking
package phase

import (
	"github.com/notional-labs/cookiemonster/transaction"
)

func MakeClaimTx(keyName string) transaction.Transaction {

	transaction := transaction.ClaimTx{KeyName: keyName}

}
