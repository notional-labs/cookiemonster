package db

import (
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

var (
	DefaultAddressToCMAddressDBDir = "/.cookiemonster/AddressToCmAddressDB.db"
	DefautlAddressToCMAddressDB    AddressToCMAddressDB
)

// // db to store user osmo address and their respective cookimonster address
// func OpenRegisteredAccountDB(dbDir string) (*leveldb.DB, error) {

// 	db, err := leveldb.OpenFile(dbDir, nil)
// 	return db, err
// }

type AddressToCMAddressDB struct {
	DB *leveldb.DB
}

func MustOpenDB(dbDir string) *leveldb.DB {
	homeDir, _ := os.UserHomeDir()
	fullDir := homeDir + dbDir
	db, err := leveldb.OpenFile(fullDir, nil)
	if err != nil {
		panic(err)
	}
	return db
}

func (d AddressToCMAddressDB) GetCMAddressForAddress(address string) (string, error) {
	addressBz, err := sdk.GetFromBech32(address, "osmo")
	if err != nil {
		return "", err
	}
	registeredAddressBz, err := d.DB.Get(addressBz, nil)
	if err == errors.ErrNotFound {
		return "", nil
	} else if err != nil {
		return "", err
	}

	registeredAddress, err := sdk.Bech32ifyAddressBytes("osmo", registeredAddressBz)
	return registeredAddress, err
}

func (d AddressToCMAddressDB) SetCMAddressForAddress(address string, registeredAccountAddress string) error {
	addressBz, err := sdk.GetFromBech32(address, "osmo")
	if err != nil {
		return err
	}
	registeredAddressBz, err := sdk.GetFromBech32(registeredAccountAddress, "osmo")
	if err != nil {
		return err
	}
	err = d.DB.Put(addressBz, registeredAddressBz, nil)
	return err
}

func (d AddressToCMAddressDB) DeleteCMAddressForAddress(address string) error {
	addressBz, err := sdk.GetFromBech32(address, "osmo")
	if err != nil {
		return err
	}
	err = d.DB.Delete(addressBz, nil)
	return err
}
