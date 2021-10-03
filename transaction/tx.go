package transaction

import "fmt"

type Tx interface {
	Execute() (string, error)
	Report()
	Prompt()
	// Type() string
}

type Txs []Tx

func HandleTx(tx Tx) error {
	tx.Prompt()
	yesOrNo := Confirmation()
	if yesOrNo == false {
		return nil
	}
	txHash, err := tx.Execute()
	if err != nil {
		return err
	}

	fmt.Printf("\nTransaction sucessful, Tx hash: %s\n", txHash)
	tx.Report()
	return nil
}

func HandleTxs(txs Txs) error {
	for _, tx := range txs {
		err := HandleTx(tx)
		if err != nil {
			return err
		}
	}
	return nil
}

func Confirmation() bool {
	fmt.Println("\nContinue [y/n] ?\n")
	var yesOrNo string
	fmt.Scanf("%s", &yesOrNo)
	if yesOrNo == "y" || yesOrNo == "yes" {
		return true
	}
	fmt.Println("Skip this transaction\n")
	return false
}
