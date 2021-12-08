package cli

import (
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/notional-labs/cookiemonster/accountmanager"
	"github.com/notional-labs/cookiemonster/osmosis"
	"github.com/spf13/cobra"
)

func AddMasterKey(privKeyBz []byte) error {
	ctx := osmosis.GetDefaultClientContext()
	kb := ctx.Keyring

	privKeyForMaster := &secp256k1.PrivKey{Key: privKeyBz}

	_, err := kb.WriteLocalKey(accountmanager.MasterKey, privKeyForMaster, hd.PubKeyType("secp256k1"))
	if err != nil {
		return err
	}
	return nil
}

func NewAddMasterKey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-master-key",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			privKeyBz, err := hex.DecodeString(args[0])
			if err != nil {
				return err
			}
			err = accountmanager.AddMasterKey(privKeyBz)
			if err != nil {
				return err
			}

			return nil
		},
	}
	return cmd
}
