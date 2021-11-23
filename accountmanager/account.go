package accountmanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/notional-labs/cookiemonster/invest"
	"github.com/notional-labs/cookiemonster/osmosis"
)

type AccountManager struct {
	Name string
	Seed string
	Num  int
}

func (am AccountManager) CreateNewAccount() string {
	ctx := osmosis.GetDefaultClientContext()
	kb := ctx.Keyring
	seed := am.Seed
	hdPath := "m/44'/118'/0'/0/"

	//  m / purpose' / coin_type' / account' / change / address_index
	// key name will be the same as address_index
	// create new account with address_index = number of keys derived from seed

	newAccountIndex := strconv.Itoa(am.Num)

	hdPath += hdPath + newAccountIndex

	keyringAlgos, _ := kb.SupportedAlgorithms()
	algoStr := "secp256k1"
	algo, err := keyring.NewSigningAlgoFromString(algoStr, keyringAlgos)
	if err != nil {
		panic(err)
	}

	k, err := kb.NewAccount(am.Name+newAccountIndex, seed, "", hdPath, algo)

	if err != nil {
		panic(err)
	}

	am.Num = am.Num + 1

	addrBz := k.GetAddress()

	return addrBz.String()
}

func (AccountManager) CreateDefautInvestmentsFromAccount() {

}

func LoadAccountManagerFromFile(fileLocation string) (*AccountManager, error) {
	file, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println("Unable to open json at " + fileLocation)
		return nil, err
	}
	reader := bufio.NewReader(file)
	jsonData, _ := ioutil.ReadAll(reader)

	var am *AccountManager
	jsonErr := json.Unmarshal(jsonData, &am)
	if jsonErr != nil {
		fmt.Println("Unable to map JSON at " + fileLocation + " to Investments")
		return nil, jsonErr
	}
	return am, nil
}

func DumpAccountManagerToFile(am *AccountManager, fileLocation string) error {

	bz, _ := json.MarshalIndent(am, "", " ")

	err := ioutil.WriteFile(fileLocation, bz, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (am AccountManager) GetDefaultInvestments() invest.Investments {

	investments := invest.Investments{}
	for i := 0; i < am.Num; i++ {
		investment := invest.Investment{
			KeyName:         am.Name + strconv.Itoa(i),
			TransferTo:      nil,
			PoolPercentage:  50,
			StakePercentage: 50,
			PoolStrategy:    invest.PoolStrategy{Name: "custom", Config: map[string]int{"1": 100}, ConfigDenom: "percentages"},
			Duration:        "14days",
			StakeAddress:    "",
		}
		investments = append(investments, investment)
	}
	return investments
}
