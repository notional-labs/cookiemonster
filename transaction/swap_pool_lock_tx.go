package transaction

// import (
// 	"fmt"
// 	"os"

// 	"gopkg.in/yaml.v3"
// )

// type SwapAndPoolAndLockTx struct {
// 	SwapAndPool SwapAndPoolTx
// 	Lock        LockTx
// }

// func (swapAndPoolAndLockTx SwapAndPoolAndLockTx) Execute() (string, error) {
// 	swapAndPoolAndLockTx.SwapAndPool.Execute()
// 	swapAndPoolAndLockTx.Lock.Execute()
// }

// func (swapAndPoolAndLockTx SwapAndPoolAndLockTx) Report() {

// 	swapAndPoolAndLockOpt := swapAndPoolAndLockTx.SwapAndPoolAndLockOpt
// 	keyName := swapAndPoolAndLockTx.KeyName

// 	f, _ := os.OpenFile("report", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

// 	f.WriteString("\nSwapAndPoolAndLock Transaction\n")
// 	f.WriteString("\nKeyname: " + keyName + "\n")
// 	f.WriteString("\nSwapAndPoolAndLock Option\n\n")

// 	txData, _ := yaml.Marshal(swapAndPoolAndLockOpt)
// 	_, _ = f.Write(txData)
// 	f.WriteString(Seperator)

// 	f.Close()
// }

// func (swapAndPoolAndLockTx SwapAndPoolAndLockTx) Prompt() {
// 	swapAndPoolAndLockOpt := swapAndPoolAndLockTx.SwapAndPoolAndLockOpt
// 	keyName := swapAndPoolAndLockTx.KeyName
// 	fmt.Print(Seperator)
// 	fmt.Print("\nSwapAndPoolAndLock Transaction\n")
// 	fmt.Print("\nKeyname: " + keyName + "\n")
// 	fmt.Print("\nSwapAndPoolAndLock Option\n\n")
// 	fmt.Printf("%+v\n", swapAndPoolAndLockOpt)

// }
