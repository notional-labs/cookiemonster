package invest

import (
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/cookiemonster/invest/phase"
	"github.com/notional-labs/cookiemonster/query"
	"github.com/notional-labs/cookiemonster/transaction"
)

type Fund struct {
	KeyName         string
	TransferTo      map[string]float32
	PoolPercentage  int
	StakePercentage int
	PoolStrategy    phase.PoolStrategy
	StakeAddress    sdk.ValAddress
}

type Funds []Fund

func (fund Fund) Invest() error {

	keyName := fund.KeyName
	poolStrategy := fund.PoolStrategy

	// 1 claim reward
	claimTx := transaction.ClaimTx{KeyName: keyName}
	// execute claim tx right away
	err := HandleTransaction(claimTx)
	if err != nil {
		return err
	}

	uosmoBalance, err := query.QueryUosmoBalance(keyName)
	if err != nil {
		return err
	}

	// 2 pool
	// caculate pool amount = pool percentage of uosmoBalance
	totalPoolAmount := XPercentageOf(uosmoBalance, fund.PoolPercentage)

	// create pooling transaction from strategy, keyname, totalpoolamount
	poolingBatch := poolStrategy.MakeTransactions(keyName, totalPoolAmount)

	for _, transaction := range poolingBatch {
		err := HandleTransaction(transaction)
		if err != nil {
			return err
		}
	}

	// 3 stake
	stakeAmount := XPercentageOf(uosmoBalance, fund.StakePercentage)
	delegateOpt := transaction.DelegateOption{
		Amount:  sdk.NewIntFromBigInt(stakeAmount),
		ValAddr: fund.StakeAddress,
		Denom:   "uosmo",
	}
	delegateTx := transaction.DelegateTx{KeyName: keyName, DelegateOpt: delegateOpt}
	err = HandleTransaction(delegateTx)
	if err != nil {
		return err
	}

	// 4 transfer

	return nil
}

func HandleTransaction(transaction transaction.Transaction) error {
	transaction.Prompt()

	transaction.Execute()

	transactionHash, err := transaction.Execute()
	if err != nil {
		return err
	}

	fmt.Printf("tx hash: %s\n", transactionHash)
	transaction.Report()
	return nil
}

func XPercentageOf(a *big.Int, x int) *big.Int {
	out := &big.Int{}
	out.Mul(a, big.NewInt(int64(x)))
	out.Div(a, big.NewInt(100))
	return out
}
