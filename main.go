package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, 世界")
	/**
	1. load wallets into memory
	2. load reward targets, dispursementStrategy per wallet
	3. for each wallet+target: evaluate shouldCollectRewards
	4. if !shouldCollectRewards, continue.
	5. if shouldCollectRewards, collect.
	6. Pass rewards to dispursementStrategy
	**/
}
