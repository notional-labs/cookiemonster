package query

import (
	"fmt"

	"github.com/notional-labs/cookiemonster/osmosis"

	"github.com/cosmos/cosmos-sdk/types"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
)

func QueryTx(txHash string) (*types.TxResponse, error) {
	clientCtx := osmosis.DefaultClientCtx

	if txHash == "" {
		return nil, fmt.Errorf("argument should be a tx hash")
	}

	// If hash is given, then query the tx by hash.
	output, err := authclient.QueryTx(clientCtx, txHash)
	if err != nil {
		return nil, err
	}

	if output.Empty() {
		return nil, fmt.Errorf("no transaction found with hash %s", txHash)
	}
	return output, nil
}
