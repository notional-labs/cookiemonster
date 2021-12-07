package db

import (
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

var (
	DefaultAddressToCMAddressDBDir = "/.cookiemonster/AddressToCMAddressDB"
	DefaultAddressToCMAddressDB    = AddressToCMAddressDB{MustOpenDB(DefaultAddressToCMAddressDBDir)}
)

// // db to store user osmo address and their respective cookimonster address
// func OpenRegisteredAccountDB(dbDir string) (*leveldb.DB, error) {

// 	db, err := leveldb.OpenFile(dbDir, nil)
// 	return db, err
// }

func MustOpenDB(dbDir string) *leveldb.DB {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	fullDir := homeDir + DefaultAddressToCMAddressDBDir
	db, err := leveldb.OpenFile(fullDir, nil)
	if err != nil {
		panic(err)
	}
	return db
}

type AddressToCMAddressDB struct {
	db *leveldb.DB
}

func (d AddressToCMAddressDB) GetCMAddressForAddress(address string) (string, error) {
	addressBz, err := sdk.GetFromBech32(address, "osmo")
	if err != nil {
		return "", err
	}
	registeredAddressBz, err := d.db.Get(addressBz, nil)
	if err == errors.ErrNotFound {
		return "", nil
	} else if err != nil {
		return "", err
	}

	registeredAddress, err := sdk.Bech32ifyAddressBytes("osmo", registeredAddressBz)
	return registeredAddress, err
}

func (d AddressToCMAddressDB) SetCMAddressForAddress(address string, cmAddress string) error {
	addressBz, err := sdk.GetFromBech32(address, "osmo")
	if err != nil {
		return err
	}
	registeredAddressBz, err := sdk.GetFromBech32(cmAddress, "osmo")
	if err != nil {
		return err
	}
	err = d.db.Put(addressBz, registeredAddressBz, nil)
	return err
}

func (d AddressToCMAddressDB) DeleteCMAccountForAddress(address string) error {
	addressBz, err := sdk.GetFromBech32(address, "osmo")
	if err != nil {
		return err
	}
	err = d.db.Delete(addressBz, nil)
	return err
}
