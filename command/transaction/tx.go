package transaction

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Tx interface {
	Execute(*cobra.Command) (string, error)
	Report(string)
	Prompt()
}

type Txs []Tx

// HandleTx print out the info of transaction, ask for permission, execute transaction
// and log to a file in reportPath
func HandleTx(tx Tx, cmd *cobra.Command, reportPath string) error {
	tx.Prompt()
	// yes := Confirmation()
	// if !yes {
	// 	return nil
	// }
	txHash, err := tx.Execute(cmd)
	if err != nil {
		return err
	}

	fmt.Printf("\nTransaction sucessful, Tx hash: %s\n", txHash)
	if reportPath != "" {
		tx.Report(reportPath)
	}
	return nil
}

func HandleTxs(txs Txs, cmd *cobra.Command, reportPath string) error {
	for _, tx := range txs {
		err := HandleTx(tx, cmd, reportPath)
		if err != nil {
			return err
		}
	}
	return nil
}

// func Confirmation() bool {
// 	fmt.Print("\nContinue [y/n] ?\n\n")
// 	var yesOrNo string
// 	fmt.Scanf("%s", &yesOrNo)
// 	if yesOrNo == "y" || yesOrNo == "yes" {
// 		return true
// 	}
// 	fmt.Print("Skip this transaction\n\n")
// 	return false
// }
