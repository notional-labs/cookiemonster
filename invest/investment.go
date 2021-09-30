package invest

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/cookiemonster/invest/pool"
	"github.com/notional-labs/cookiemonster/query"
	"github.com/notional-labs/cookiemonster/transaction"
)

type Investment struct {
	KeyName         string
	TransferTo      map[string]float32
	PoolPercentage  int
	StakePercentage int
	PoolStrategy    pool.PoolStrategy
	StakeAddress    sdk.ValAddress
}

type Investments []Investment

func (investment Investment) Invest() error {

	keyName := investment.KeyName
	poolStrategy := investment.PoolStrategy

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
	totalPoolAmount := XPercentageOf(uosmoBalance, investment.PoolPercentage)

	// create pooling transaction from strategy, keyname, totalpoolamount
	poolingBatch := poolStrategy.MakeTransactions(keyName, totalPoolAmount)

	for _, transaction := range poolingBatch {
		err := HandleTransaction(transaction)
		if err != nil {
			return err
		}
	}

	// 3 stake
	stakeAmount := XPercentageOf(uosmoBalance, investment.StakePercentage)
	delegateOpt := transaction.DelegateOption{
		Amount:  sdk.NewIntFromBigInt(stakeAmount),
		ValAddr: investment.StakeAddress,
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

// Cal x percent of a
func XPercentageOf(a *big.Int, x int) *big.Int {
	out := &big.Int{}
	out.Mul(a, big.NewInt(int64(x)))
	out.Div(a, big.NewInt(100))
	return out
}

func LoadInvestmentsFromFile(fileLocation string) ([]Investment, error) {
	file, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println("Unable to open json at " + fileLocation)
		return nil, err
	}
	reader := bufio.NewReader(file)
	jsonData, _ := ioutil.ReadAll(reader)

	var investments []Investment
	jsonErr := json.Unmarshal(jsonData, &investments)
	if jsonErr != nil {
		fmt.Println("Unable to map JSON at " + fileLocation + " to Invesments")
		return nil, jsonErr
	}
	return investments, nil
}
