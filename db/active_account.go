package db

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/syndtr/goleveldb/leveldb/errors"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	defaultActiveDBDir = "/.cookiemonster/active_accounts.db"
)

// db to store user osmo address and their respective cookimonster address
func OpenActiveAccountDB(dbDir string) (*leveldb.DB, error) {
	db, err := leveldb.OpenFile(dbDir, nil)
	return db, err
}

func CheckIfAccountAddressIsActive(db *leveldb.DB, accountAddress string) (string, error) {
	addressBz, err := sdk.GetFromBech32(accountAddress, "osmo")
	if err != nil {
		return "unknown", err
	}
	_, err = db.Get(addressBz, nil)
	if err == errors.ErrNotFound {
		return "inactive", nil
	} else if err != nil {
		return "unknown", err
	}

	return "active", nil
}

func SetActiveAccountAddress(db *leveldb.DB, accountAddress string) error {
	addressBz, err := sdk.GetFromBech32(accountAddress, "osmo")
	if err != nil {
		return err
	}
	err = db.Put(addressBz, []byte{1}, nil)
	return err
}

func DeleteActiveAccountAddress(db *leveldb.DB, accountAddress string) error {
	addressBz, err := sdk.GetFromBech32(accountAddress, "osmo")
	if err != nil {
		return err
	}
	err = db.Delete(addressBz, nil)
	return err
}
