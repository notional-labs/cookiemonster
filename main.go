package main

import (
	"fmt"
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
	fmt.Println(10e2)

	err1 := fmt.Errorf("test").Error()
	err2 := fmt.Errorf("test").Error()
	fmt.Println(err1 == err2)
	// txf := transaction.NewTxFactoryFromClientCtx(osmosis.DefaultClientCtx)
	// fmt.Printf("%+v\n", txf)
}
