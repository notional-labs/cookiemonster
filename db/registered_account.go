package db

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

var (
	defaultRegisteredAccountDBDir = "/.cookiemonster"
)

// db to store user osmo address and their respective cookimonster address
func OpenRegisteredAccountDB(dbDir string) (*leveldb.DB, error) {

	db, err := leveldb.OpenFile(dbDir, nil)
	return db, err
}

func GetRegisterAddressForAddress(db *leveldb.DB, address string) (string, error) {
	addressBz, err := sdk.GetFromBech32(address, "osmo")
	if err != nil {
		return "", err
	}
	registeredAddressBz, err := db.Get(addressBz, nil)
	if err == errors.ErrNotFound {
		return "", nil
	} else if err != nil {
		return "", err
	}

	registeredAddress, err := sdk.Bech32ifyAddressBytes("osmo", registeredAddressBz)
	return registeredAddress, err
}

func SetRegisterAddressBzForAddress(db *leveldb.DB, address string, registeredAccountAddress string) error {
	addressBz, err := sdk.GetFromBech32(address, "osmo")
	if err != nil {
		return err
	}
	registeredAddressBz, err := sdk.GetFromBech32(registeredAccountAddress, "osmo")
	if err != nil {
		return err
	}
	err = db.Put(addressBz, registeredAddressBz, nil)
	return err
}
