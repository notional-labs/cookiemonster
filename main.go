package main

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/cookiemonster/invest"
	"github.com/notional-labs/cookiemonster/query"
	"github.com/notional-labs/cookiemonster/transaction"
	//
	// "github.com/notional-labs/cookiemonster/command/transaction"
)

func main() {
	// toAddr, _ := sdk.AccAddressFromBech32("osmo13k9w2pexxtyyfuw7fmxh6rpwwl9udxkk26nfle")
	// bankSendOpt := transaction.BankSendOption{
	// 	ToAddr: toAddr,
	// 	Denom:  "stake",
	// 	Amount: sdk.NewInt(12),
	// }
	// fmt.Println(transaction.BankSend("april", bankSendOpt))
	// pool, err := query.QuerySpotPrice()
	// if err != nil {
	// 	fmt.Println(err, "fsasdfa")
	// }
	// fmt.Println(10e2)
	// confirmation := ""
	// fmt.Scanln(confirmation)
	// fmt.Println(confirmation)
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// text, _ := reader.ReadString('\n')
	// fmt.Println(text)
	// fmt.Println("Enter text: ")
	// text2 := ""
	// fmt.Scanln(&text2)
	// fmt.Println(text2)

	// ln := ""
	// fmt.Scanf("%s", &ln)
	// tx := transaction.ClaimTx{KeyName: "koule"}
	// txhash, err := tx.Execute()
	// fmt.Println(txhash)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	/*
		userHomeDir, _ := os.UserHomeDir()
		investments, _ := invest.LoadInvestmentsFromFile(userHomeDir + "/auto-farm/investments.json")
		investment := investments[0]
		strategy := investment.PoolStrategy

		// uosmoBalance, _ := query.QueryUosmoBalance("koule")
		// if err != nil {
		// 	return err
		// }

		// 2 pool
		// caculate pool amount = pool percentage of uosmoBalance
		totalPoolAmount := invest.XPercentageOf(big.NewInt(15372), investment.PoolPercentage)
		fmt.Println(totalPoolAmount)

		M := invest.MakeSwapAndPoolTxs("koule", totalPoolAmount, strategy)
		for _, i := range M {
			_, err := i.Execute()
			if err != nil {
				fmt.Println(err)
			}
			i.Prompt()
			// i.Report()
		}
	*/

	//============ Staking ===============
	// Input
	stakeAddress := ""
	keyName := ""
	stakePercentage := 95

	investment := invest.Investment{
		StakePercentage: stakePercentage,
	}
	uosmoBalance, err := query.QueryUosmoBalance(keyName)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	stakeAmount := invest.XPercentageOf(uosmoBalance, investment.StakePercentage)

	valAddress, err := sdk.ValAddressFromBech32(stakeAddress)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	delegateOpt := transaction.DelegateOption{
		Amount:  sdk.NewIntFromBigInt(stakeAmount),
		ValAddr: valAddress,
		Denom:   "uosmo",
	}

	delegateTx := transaction.DelegateTx{KeyName: keyName, DelegateOpt: delegateOpt}
	res, err := delegateTx.Execute()
	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Println((res))

	//fmt.Printf("%+v\n", investment)
	// err := investment[0].Invest()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// txf := transaction.NewTxFactoryFromClientCtx(osmosis.DefaultClientCtx)

}
