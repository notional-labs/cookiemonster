package db

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/syndtr/goleveldb/leveldb/errors"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	DefaultAddressToCMKeyNameDBDir = "/.cookiemonster/AddressToCMKeyName.db"
	DefaultAddressToCMKeyNameDB    = AddressToCMKeyDB{MustOpenDB(DefaultAddressToCMKeyNameDBDir)}
)

type AddressToCMKeyDB struct {
	db *leveldb.DB
}

func (d AddressToCMKeyDB) GetCMKeyNameForAddress(address string) (string, error) {
	addressBz, err := sdk.GetFromBech32(address, "osmo")
	if err != nil {
		return "", err
	}
	keyNameBz, err := d.db.Get(addressBz, nil)
	if err == errors.ErrNotFound {
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
	err = d.db.Put(addressBz, []byte(keyName), nil)
	return err
}

func (d AddressToCMKeyDB) DeleteCMKeyNameForAddress(address string) error {
	addressBz, err := sdk.GetFromBech32(address, "osmo")
	if err != nil {
		return err
	}
	err = d.db.Delete(addressBz, nil)
	return err
}
