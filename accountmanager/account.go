package accountmanager

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/cookiemonster/osmosis"
)

var (
	defaultAccountManagerFile = "accountmanager"
	DefaultAccountManager     = MustLoadAccountManagerFromFile(defaultAccountManagerFile)
)

type AccountManager struct {
	MasterKey        []byte
	NumOfAccount     int
	Passphrase       string
	HashedPassphrase string
}

func (am *AccountManager) CreateNewPrivKeyForAddress(Address string) (cryptotypes.PrivKey, error) {
	// ctx := osmosis.GetDefaultClientContext()
	// kb := ctx.Keyring

	masterKey := am.MasterKey
	toBeHashed := append(masterKey, []byte(Address)...)
	privKeyBz32ForAddress := sha256.Sum256(toBeHashed)

	privKeyBzForAddress := privKeyBz32ForAddress[:]

	// accountIdString := am.HashedPassphrase + "_" + strconv.Itoa(am.NumOfAccount)
	privKeyForAddress, err := legacy.PrivKeyFromBytes(privKeyBzForAddress)

	// uid := "acc" + "-" + accountIdString

	return privKeyForAddress, err
	// kb.WriteLocalKey(uid, privKeyForAddress, hd.PubKeyType("secp256k1"))

}

// import
func (am *AccountManager) RegisterAccountForAddress(Address string) (sdk.AccAddress, error) {
	ctx := osmosis.GetDefaultClientContext()
	kb := ctx.Keyring

	privKeyForAddress, err := am.CreateNewPrivKeyForAddress(Address)
	if err != nil {
		return nil, err
	}

	accountIdString := am.HashedPassphrase + "_" + strconv.Itoa(am.NumOfAccount)
	uid := "acc" + "-" + accountIdString
	_, err = kb.WriteLocalKey(uid, privKeyForAddress, hd.PubKeyType("secp256k1"))
	if err != nil {
		return nil, err
	}
	generatedAddress := sdk.AccAddress(privKeyForAddress.PubKey().Address().Bytes())
	return generatedAddress, nil
}

// func (AccountManager) CreateDefautInvestmentsFromAccount() {

// }

func MustLoadAccountManagerFromFile(fileDir string) *AccountManager {
	file, err := os.Open(fileDir)
	if err != nil {
		fmt.Println("Unable to open json at " + fileDir)
		panic(err)
	}
	reader := bufio.NewReader(file)
	jsonData, _ := ioutil.ReadAll(reader)

	var am *AccountManager
	jsonErr := json.Unmarshal(jsonData, &am)
	if jsonErr != nil {
		fmt.Println("Unable to map JSON at " + fileDir + " to Investments")
		panic(err)
	}
	return am
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

func CreateAccount(am AccountManager)

func DumpAccountManagerToFile(am *AccountManager, fileLocation string) error {

	bz, _ := json.MarshalIndent(am, "", " ")

	err := ioutil.WriteFile(fileLocation, bz, 0644)
	if err != nil {
		return err
	}
	return nil
}

// func (am AccountManager) GetDefaultInvestments() invest.Investments {

// 	investments := invest.Investments{}
// 	for i := 0; i < am.Num; i++ {
// 		investment := invest.Investment{
// 			KeyName:         am.Name + strconv.Itoa(i),
// 			TransferTo:      nil,
// 			PoolPercentage:  50,
// 			StakePercentage: 50,
// 			PoolStrategy:    invest.PoolStrategy{Name: "custom", Config: map[string]int{"1": 100}, ConfigDenom: "percentages"},
// 			Duration:        "14days",
// 			StakeAddress:    "",
// 		}
// 		investments = append(investments, investment)
// 	}
// 	return investments
// }
