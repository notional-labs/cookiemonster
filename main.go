package main

import (
	"fmt"
	"time"
	// "github.com/notional-labs/cookiemonster/osmosis"
	// "github.com/notional-labs/cookiemonster/transaction"
)

// func main() {
// 	// toAddr, _ := sdk.AccAddressFromBech32("osmo13k9w2pexxtyyfuw7fmxh6rpwwl9udxkk26nfle")
// 	// bankSendOpt := transaction.BankSendOption{
// 	// 	ToAddr: toAddr,
// 	// 	Denom:  "stake",
// 	// 	Amount: sdk.NewInt(12),
// 	// }
// 	// fmt.Println(transaction.BankSend("april", bankSendOpt))
// 	// pool, err := query.QuerySpotPrice()
// 	// if err != nil {
// 	// 	fmt.Println(err, "fsasdfa")
// 	// }
// 	// fmt.Println(10e2)
// 	// confirmation := ""
// 	// fmt.Scanln(confirmation)
// 	// fmt.Println(confirmation)
// 	// reader := bufio.NewReader(os.Stdin)
// 	// fmt.Print("Enter text: ")
// 	// text, _ := reader.ReadString('\n')
// 	// fmt.Println(text)
// 	// fmt.Println("Enter text: ")
// 	// text2 := ""
// 	// fmt.Scanln(&text2)
// 	// fmt.Println(text2)

// 	// ln := ""
// 	// fmt.Scanf("%s", &ln)
// 	// tx := transaction.ClaimTx{KeyName: "koule"}
// 	// txhash, err := tx.Execute()
// 	// fmt.Println(txhash)
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// }
// 	investments, _ := invest.LoadInvestmentsFromFile("/home/pegasus/auto-farm/investments.json")
// 	investment := investments[0]
// 	strategy := investment.PoolStrategy

// 	// uosmoBalance, _ := query.QueryUosmoBalance("koule")
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// 2 pool
// 	// caculate pool amount = pool percentage of uosmoBalance
// 	totalPoolAmount := invest.XPercentageOf(big.NewInt(15372), investment.PoolPercentage)
// 	fmt.Println(totalPoolAmount)

// 	M := strategy.MakeTransactions("koule", totalPoolAmount)
// 	M[1].Prompt()

// 	fmt.Printf("%+v\n", investment)
// 	// err := investment[0].Invest()
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// }
// 	// txf := transaction.NewTxFactoryFromClientCtx(osmosis.DefaultClientCtx)

// }
func DoneAsync() chan int {
	r := make(chan int)
	fmt.Println("Warming up ...")
	go func() {
		time.Sleep(3 * time.Second)
		r <- 1
		fmt.Println("Done ...")
	}()
	return r
}

func main() {
	fmt.Println("Let's start ...")
	val := DoneAsync()
	fmt.Println("Done is running ...")
	fmt.Println(<-val)
}
