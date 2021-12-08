package db

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/syndtr/goleveldb/leveldb/errors"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	DefaultAddressToCMKeyNameDBDir = "/.cookiemonster/AddressToCMKeyName.db"
	DefaultAddressToCMKeyNameDB    AddressToCMKeyDB
)

type AddressToCMKeyDB struct {
	DB *leveldb.DB
}

func (d AddressToCMKeyDB) GetCMKeyNameForAddress(address string) (string, error) {
	addressBz, err := sdk.GetFromBech32(address, "osmo")
	if err != nil {
		return "", err
	}
	keyNameBz, err := d.DB.Get(addressBz, nil)
	if err == errors.ErrNotFound {
		fmt.Println("not found")
		return "", nil
	} else if err != nil {
		return "", err
	}

	return string(keyNameBz), nil
}

func (d AddressToCMKeyDB) SetCMKeyNameForAddress(address string, keyName string) error {
	addressBz, err := sdk.GetFromBech32(address, "osmo")
	if err != nil {
		return err
	}
	err = d.DB.Put(addressBz, []byte(keyName), nil)
	return err
}

func (d AddressToCMKeyDB) DeleteCMKeyNameForAddress(address string) error {
	addressBz, err := sdk.GetFromBech32(address, "osmo")
	if err != nil {
		return err
	}
	err = d.DB.Delete(addressBz, nil)
	return err
}
