package main

import (
	"fmt"
	"os"
	// "github.com/notional-labs/cookiemonster/osmosis"
	// "github.com/notional-labs/cookiemonster/transaction"
)

func main() {
	s := "kkkkkkk"

	b := []byte(s)

	fmt.Println(string(b))

	homeDir, err := os.UserHomeDir()
	if err == nil {
		fmt.Println(homeDir)
	}

}
