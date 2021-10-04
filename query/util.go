package query

import (
	"github.com/cosmos/cosmos-sdk/client"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// set fields in client context with keyname
func GetAddressFromKey(clientCtx client.Context, keyName string) (sdk.AccAddress, error) {
	from := keyName
	fromAddr, _, _, err := client.GetFromFields(clientCtx.Keyring, from, clientCtx.GenerateOnly)
	if err != nil {
		return nil, err
	}

	return fromAddr, nil
}
