package transaction

import (
	"fmt"
	"os"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/notional-labs/cookiemonster/command/query"
	"github.com/osmosis-labs/osmosis/x/lockup/types"
	"gopkg.in/yaml.v3"
)

type LockOption struct {
	Duration string
	Amount   sdk.Int
	Denom    string
}

func Lock(cmd *cobra.Command, keyName string, lockOpt LockOption, gas uint64) (string, error) {
	clientCtx := client.GetClientContextFromCmd(cmd)
	clientCtx, err := SetKeyNameToContext(clientCtx, keyName)
	if err != nil {
		return "", err
	}
	txf := NewTxFactoryFromClientCtx(clientCtx).WithGas(gas)
	var durationString string

	if lockOpt.Duration == "14days" {
		durationString = "1209600s"
	} else if lockOpt.Duration == "7days" {
		durationString = "604800s"
	} else if lockOpt.Duration == "1day" {
		durationString = "86400s"
	} else {
		return "", fmt.Errorf("unknown duration (bonding period)")
	}

	duration, err := time.ParseDuration(durationString)
	if err != nil {
		return "", err
	}

	msg := types.NewMsgLockTokens(
		clientCtx.GetFromAddress(),
		duration,
		sdk.Coins{{Denom: lockOpt.Denom, Amount: lockOpt.Amount}},
	)

	code, txHash, err := BroadcastTx(clientCtx, txf, msg)
	if err != nil {
		return txHash, err
	}
	if code != 0 {
		return txHash, fmt.Errorf("tx failed with code %d", code)
	}
	broadcastedTx, err := query.QueryTxWithRetry(cmd, txHash, 4)
	if err != nil {
		return txHash, err
	}
	if broadcastedTx.Code == 11 {
		return txHash, fmt.Errorf("insufficient fee")

	}
	if broadcastedTx.Code != 0 {
		return txHash, fmt.Errorf("tx failed with code %d", code)
	}
	return txHash, nil

}

type LockTx struct {
	KeyName string
	LockOpt LockOption
	Hash    string
}

func (lockTx LockTx) Execute(cmd *cobra.Command) (string, error) {

	keyName := lockTx.KeyName
	lockOpt := lockTx.LockOpt
	gas := 200000000
	var err error
	var txHash string

	// if tx failed because of insufficient fee , retry
	for i := 0; i < 4; i++ {
		fmt.Println("\n---------------")
		fmt.Printf("\n Try %d times\n\n", i+1)
		txHash, err = Lock(cmd, keyName, lockOpt, uint64(gas))

		if err == nil {
			lockTx.Hash = txHash
			return txHash, nil
		}
		if err.Error() == "insufficient fee" {
			fmt.Println("\nTx failed because of insufficient fee, try again with higher gas\n")
			gas += 300000
		} else {
			fmt.Println("\n" + err.Error() + " try again\n")
		}
	}
	return txHash, err
}

func (lockTx LockTx) Report(reportPath string) {

	lockOpt := lockTx.LockOpt
	keyName := lockTx.KeyName

	f, _ := os.OpenFile(reportPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	f.WriteString("\nLock Transaction\n")
	f.WriteString("\nKeyname: " + keyName + "\n")
	f.WriteString("\nLock Option\n\n")

	txData, _ := yaml.Marshal(lockOpt)
	_, _ = f.Write(txData)
	f.WriteString("\ntx hash: " + lockTx.Hash + "\n")
	f.WriteString(Seperator)

	f.Close()
}

func (lockTx LockTx) Prompt() {
	lockOpt := lockTx.LockOpt
	keyName := lockTx.KeyName
	fmt.Print(Seperator)
	fmt.Print("\nLock Transaction\n")
	fmt.Print("\nKeyname: " + keyName + "\n")
	fmt.Print("\nLock Option\n\n")
	fmt.Printf("%+v\n", lockOpt)

}
