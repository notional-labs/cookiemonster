package main

import (
	"os"

	"github.com/notional-labs/cookiemonster/accountmanager"
	"github.com/notional-labs/cookiemonster/invest"
	//
	// "github.com/notional-labs/cookiemonster/transaction"
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
	// userHomeDir, _ := os.UserHomeDir()
	// investments, _ := invest.LoadInvestmentsFromFile(userHomeDir + "/auto-farm/investments.json")
	// investment := investments[0]
	// strategy := investment.PoolStrategy

	// uosmoBalance, _ := query.QueryUosmoBalance("koule")
	// if err != nil {
	// 	return err
	// }

	// 2 pool
	// caculate pool amount = pool percentage of uosmoBalance
	// totalPoolAmount := invest.XPercentageOf(big.NewInt(15372), investment.PoolPercentage)
	// fmt.Println(totalPoolAmount)

	// M := invest.MakeSwapAndPoolTxs("koule", totalPoolAmount, strategy)
	// for _, i := range M {
	// 	_, err := i.Execute()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	i.Prompt()
	// 	// i.Report()
	// }

	// fmt.Printf("%+v\n", investment)
	// err := investment[0].Invest()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// txf := transaction.NewTxFactoryFromClientCtx(osmosis.DefaultClientCtx)
	userHomeDir, _ := os.UserHomeDir()
	am, err := accountmanager.LoadAccountManagerFromFile(userHomeDir + "/accountmanager.json")
	if err != nil {
		panic(err)
	}
	//	addr := am.CreateNewAccount()
	//	if err != nil {
	//		panic(err)
	//	}
	accountmanager.DumpAccountManagerToFile(am, userHomeDir+"/accountmanager.json")
	investments := am.GetDefaultInvestments()
	invest.DumpInvestmentsToFile(userHomeDir+"/cookiemonster/investments.json", investments)
	//	fmt.Println(addr)

}
