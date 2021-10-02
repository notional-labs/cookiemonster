package transaction

import "fmt"

type Tx interface {
	Execute() (string, error)
	Report()
	Prompt()
	// Type() string
}

type Txs []Tx

func HandleTransaction(tx Tx) error {
	tx.Prompt()

	tx.Execute()

	txHash, err := tx.Execute()
	if err != nil {
		return err
	}

	fmt.Printf("tx hash: %s\n", txHash)
	tx.Report()
	return nil
}
