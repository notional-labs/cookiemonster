package query

import (
	"fmt"
	"time"

	"github.com/notional-labs/cookiemonster/osmosis"

	"github.com/cosmos/cosmos-sdk/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
)

func QueryTx(txHash string) (*types.TxResponse, error) {
	clientCtx := osmosis.GetDefaultClientContext()

	if txHash == "" {
		return nil, fmt.Errorf("argument should be a tx hash")
	}

	// If hash is given, then query the tx by hash.
	output, err := authtx.QueryTx(clientCtx, txHash)

	if err != nil {
		return nil, err
	}

	if output.Empty() {
		return nil, fmt.Errorf("no transaction found with hash %s", txHash)
	}
	return output, nil
}

func QueryTxWithRetry(txHash string, trials int) (*types.TxResponse, error) {
	var broadcastedTx *types.TxResponse
	var err error

	for i := 0; i < trials; i++ {
		broadcastedTx, err = QueryTx(txHash)
		if err == nil {
			break
		}
		time.Sleep(3000)
	}

	return broadcastedTx, err
}
