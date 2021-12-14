package invest

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/notional-labs/cookiemonster/query"
	"github.com/notional-labs/cookiemonster/transaction"
)

type Investment struct {
	KeyName         string
	TransferTo      map[string]float32
	PoolPercentage  int
	StakePercentage int
	PoolStrategy    PoolStrategy
	Duration        string
	StakeAddress    string
}

type Investments []Investment

func (investment Investment) Invest(reportPath string) error {

	keyName := investment.KeyName

	// 1 claim reward
	claimTx := transaction.ClaimTx{KeyName: keyName}
	// execute claim tx right away
	err := transaction.HandleTx(claimTx, reportPath)
	if err != nil {
		return err
	}

	uosmoBalance, err := query.QueryUosmoBalance(keyName)
	if err != nil {
		return err
	}

	// poling
	poolStrategy := investment.PoolStrategy
	totalPoolAmount := XPercentageOf(uosmoBalance, investment.PoolPercentage)
	err = BatchPool(keyName, totalPoolAmount, poolStrategy, investment.Duration, reportPath)
	if err != nil {
		return err
	}

	// staking
	stakeAmount := XPercentageOf(uosmoBalance, investment.StakePercentage)
	err = Stake(keyName, stakeAmount, investment.StakeAddress, reportPath)
	if err != nil {
		return err
	}

	// 4 transfer

	return nil
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
		fmt.Println("Unable to map JSON at " + fileLocation + " to Investments")
		return nil, jsonErr
	}
	return investments, nil
}

func BatchPool(keyName string, totalPoolAmount *big.Int, poolStrategy PoolStrategy, duration string, reportPath string) error {
	// fmt.Println(transaction.Seperator)
	// fmt.Println("\nPooling:")
	// 2 pool
	// caculate pool amount = pool percentage of uosmoBalance

	// fmt.Println("\nTotal Pool Amount: " + totalPoolAmount.String() + "uosmo\n")
	// create pooling transaction from strategy, keyname, totalpoolamount
	swapAndPoolTxs := MakeSwapAndPoolTxs(keyName, totalPoolAmount, poolStrategy)

	err := transaction.HandleTxs(swapAndPoolTxs, reportPath)
	if err != nil {
		return err
	}
	lockTxs, err := MakeLockTxs(keyName, duration)

	if err != nil {
		return err
	}
	err = transaction.HandleTxs(lockTxs, reportPath)
	if err != nil {
		return err
	}
	return nil
}

func Stake(keyName string, stakeAmount *big.Int, stakeAddress string, reportPath string) error {
	if stakeAddress != "" {
		valAddress, err := sdk.ValAddressFromBech32(stakeAddress)
		if err != nil {
			return err
		}
		// 3 stake

		delegateOpt := transaction.DelegateOption{
			Amount:  sdk.NewIntFromBigInt(stakeAmount),
			ValAddr: valAddress,
			Denom:   "uosmo",
		}
		delegateTx := transaction.DelegateTx{KeyName: keyName, DelegateOpt: delegateOpt}
		err = transaction.HandleTx(delegateTx, reportPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (investment Investment) InvestToDie(reportPath string) error {
	keyName := investment.KeyName
	currentEpoch, err := getEpochIdx()
	if err != nil {
		return err
	}

	// 1 claim reward
	claimTx := transaction.ClaimTx{KeyName: keyName}
	// execute claim tx right away
	err = transaction.HandleTx(claimTx, reportPath)
	if err != nil {
		return err
	}
	for {
		osmosisEpoch, err := getEpochIdx()
		if err != nil {
			return err
		}

		if currentEpoch < osmosisEpoch {
			keyName := investment.KeyName
			currentEpoch, err := getEpochIdx()
			if err != nil {
				return err
			}
			// 1 claim reward
			claimTx := transaction.ClaimTx{KeyName: keyName}
			// execute claim tx right away
			err = transaction.HandleTx(claimTx, reportPath)
			if err != nil {
				return err
			}

			fmt.Println("Get Reward in Epoch ", currentEpoch)
			fmt.Println("Farming Continue ")

			break
		}
		uosmoBalance, err := query.QueryUosmoBalance(keyName)
		if err != nil {
			return err
		}

		// poling
		poolStrategy := investment.PoolStrategy
		totalPoolAmount := XPercentageOf(uosmoBalance, investment.PoolPercentage)
		err = BatchPool(keyName, totalPoolAmount, poolStrategy, investment.Duration, reportPath)
		if err != nil {
			return err
		}

		// staking
		stakeAmount := XPercentageOf(uosmoBalance, investment.StakePercentage)
		err = Stake(keyName, stakeAmount, investment.StakeAddress, reportPath)
		if err != nil {
			return err
		}

		fmt.Println("sleeping till next epoch")
		time.Sleep(1 * time.Hour)
		// 4 transfer
	}
	return nil
}

func getEpochIdx() (int64, error) {
	epoch, err := query.QueryEpoch()
	if err != nil {
		fmt.Println("Err Epoch", err)
		return -1, err
	}
	fmt.Println("Farming in epoch ", epoch)
	return epoch, nil
}
