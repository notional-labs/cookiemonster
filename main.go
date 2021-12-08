package main

import (
	"fmt"

	// "github.com/notional-labs/cookiemonster/osmosis"
	// "github.com/notional-labs/cookiemonster/transaction"
	crypto "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

func main() {
	privKeyBzForAddress := []byte{0, 1}

	// // accountIdString := am.HashedPassphrase + "_" + strconv.Itoa(am.NumOfAccount)
	// privKeyForAddress, err := legacy.PrivKeyFromBytes(privKeyBzForAddress)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// sdkAcc := sdk.AccAddress(privKeyForAddress.PubKey().Address())

	// fmt.Println(sdkAcc.String())
	// fmt.Println(9)

	keys := crypto.PrivKey{
		Key: privKeyBzForAddress,
	}
	fmt.Println(9)

	fmt.Println(keys.PubKey().Address().String())
}
