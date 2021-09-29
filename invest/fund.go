package invest

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/cookiemonster/invest/phase"
	"github.com/notional-labs/cookiemonster/transaction"
)

type Fund struct {
	KeyName         string
	TransferTo      map[string]float32
	PoolPercentage  float32
	StakePercentage float32
	PoolStrategy    phase.PoolStrategy
	StakeAddress    sdk.ValAddress
}

type Funds []Fund

func (fund Fund) Invest() {
	transactionBatch := transaction.Transactions{}

	keyName := fund.KeyName
	poolStrategy := fund.PoolStrategy

	// 1 claim reward
	claimTx := transaction.ClaimTx{KeyName: keyName}
	claimTx.Execute()

	// 2 pool
	totalPoolAmount := big.Int{}

	poolingBatch := poolStrategy.MakeTransactions(keyName)
	transactionBatch = append(transactionBatch, poolingBatch...)

	// 3 stake
	delegateOpt := transaction.DelegateOption{
		ValAddr: fund.StakeAddress,
		Denom:   fund,
	}
	delegateTx := transaction.DelegateTx{}

}
