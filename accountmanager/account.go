package accountmanager

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/cookiemonster/db"
	"github.com/notional-labs/cookiemonster/osmosis"
)

var (
	DefaultAccountManagerFile = "/.cookiemonster/accountmanager.json"
	DefaultAccountManager     AccountManager
	MasterKey                 = "master"
)

type AccountManager struct {
	MasterKey     []byte
	NumOfAccount  int
	MasterAddress string
}

type AccountManagerLoader struct {
	MasterKeyHex string
	NumOfAccount int
}

func (am *AccountManager) CreateNewPrivKeyForAddress(Address string) cryptotypes.PrivKey {
	// ctx := osmosis.GetDefaultClientContext()
	// kb := ctx.Keyring

	masterKey := am.MasterKey
	toBeHashed := append(masterKey, []byte(Address)...)
	privKeyBz32ForAddress := sha256.Sum256(toBeHashed)

	privKeyBzForAddress := privKeyBz32ForAddress[:]

	secp256k1Key := secp256k1.PrivKey{Key: privKeyBzForAddress}

	// accountIdString := am.HashedPassphrase + "_" + strconv.Itoa(am.NumOfAccount)
	// privKeyForAddress, err := legacy.PrivKeyFromBytes(privKeyBzForAddress)

	// uid := "acc" + "-" + accountIdString

	return &secp256k1Key
	// kb.WriteLocalKey(uid, privKeyForAddress, hd.PubKeyType("secp256k1"))

}

func AddMasterKey(privKeyBz []byte) error {
	ctx := osmosis.GetDefaultClientContext()
	kb := ctx.Keyring

	privKeyForMaster := &secp256k1.PrivKey{Key: privKeyBz}

	_, err := kb.WriteLocalKey(MasterKey, privKeyForMaster, hd.PubKeyType("secp256k1"))
	if err != nil {
		return err
	}
	return nil
}

func (am *AccountManager) LoadMasterKey() {
	ctx := osmosis.GetDefaultClientContext()
	kb := ctx.Keyring

	privKeyForMaster := &secp256k1.PrivKey{Key: am.MasterKey}

	masterAddressBz := sdk.AccAddress(privKeyForMaster.PubKey().Address())

	masterAddress := masterAddressBz.String()

	am.MasterAddress = masterAddress

	err := kb.Delete("master")
	if err != nil {
		fmt.Println(err)
	}

	// in regular usage this should never print a private key to the consiole.
	keyring, err := kb.WriteLocalKey(MasterKey, privKeyForMaster, hd.PubKeyType("secp256k1"))
	if err != nil {
		fmt.Println(err, keyring)
	}

	err = AddMasterKey(am.MasterKey)
	if err != nil {
		fmt.Println(err)
	}
}

// import
func (am *AccountManager) RegisterAccountForAddress(Address string) (sdk.AccAddress, error) {
	ctx := osmosis.GetDefaultClientContext()
	kb := ctx.Keyring

	privKeyForAddress := am.CreateNewPrivKeyForAddress(Address)

	accountIdString := strconv.Itoa(am.NumOfAccount)

	uid := "acc" + "-" + accountIdString
	_, err := kb.WriteLocalKey(uid, privKeyForAddress, hd.PubKeyType("secp256k1"))
	if err != nil {
		return nil, err
	}
	am.NumOfAccount += 1

	cmAddress := sdk.AccAddress(privKeyForAddress.PubKey().Address().Bytes())

	addressToCMKeyDB := db.DefaultAddressToCMKeyNameDB
	addressToCMAddressDB := db.DefautlAddressToCMAddressDB

	err = addressToCMAddressDB.SetCMAddressForAddress(Address, cmAddress.String())
	if err != nil {
		panic(err)
	}

	err = addressToCMKeyDB.SetCMKeyNameForAddress(Address, uid)
	if err != nil {
		panic(err)
	}
	err = DumpAccountManagerToFile(am, DefaultAccountManagerFile)
	if err != nil {
		panic(err)
	}
	return cmAddress, nil
}

// func (AccountManager) CreateDefautInvestmentsFromAccount() {

// }

func MustLoadAccountManagerFromFile(fileDir string) *AccountManager {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	fullDir := homeDir + fileDir

	file, err := os.Open(fullDir)
	if err != nil {
		fmt.Println("Unable to open json at " + fileDir)
		panic(err)
	}
	reader := bufio.NewReader(file)
	jsonData, _ := ioutil.ReadAll(reader)

	var aml *AccountManagerLoader
	jsonErr := json.Unmarshal(jsonData, &aml)
	if jsonErr != nil {
		fmt.Println("Unable to map JSON at " + fileDir + " to Investments")
		panic(err)
	}

	masterKey, err := hex.DecodeString(aml.MasterKeyHex)
	if err != nil {
		panic(err)
	}
	am := &AccountManager{
		MasterKey:    masterKey,
		NumOfAccount: aml.NumOfAccount,
	}

	am.LoadMasterKey()
	fmt.Println(am)

	return am
}

// I love Ngan!!!!!!
// func (am *AccountManager)LoadMasterKey() {
// 	ctx := osmosis.GetDefaultClientContext()
// 	kb := ctx.Keyring

// 	privKeyForMaster := &secp256k1.PrivKey{Key: privKeyBz}

// 	_, err := kb.WriteLocalKey("master", privKeyForMaster, hd.PubKeyType("secp256k1"))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

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
	homeDir, _ := os.UserHomeDir()

	fullDir := homeDir + fileLocation

	masterKeyHex := hex.EncodeToString(am.MasterKey)
	aml := AccountManagerLoader{
		MasterKeyHex: masterKeyHex,
		NumOfAccount: am.NumOfAccount,
	}
	bz, _ := json.MarshalIndent(aml, "", " ")
	err := ioutil.WriteFile(fullDir, bz, 0644)
	if err != nil {
		return err
	}
	return nil
}

/*
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
*/
