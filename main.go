package main

import (
	"fmt"

	"github.com/notional-labs/cookiemonster/transaction"
	// "github.com/notional-labs/cookiemonster/osmosis"
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
	a := transaction.BankSendTx{}
	fmt.Println(a.Hash)
	// txf := transaction.NewTxFactoryFromClientCtx(osmosis.DefaultClientCtx)
	// fmt.Printf("%+v\n", txf)
}
