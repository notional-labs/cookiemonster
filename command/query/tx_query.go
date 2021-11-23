package query

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/types"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
)

func QueryTx(cmd *cobra.Command, txHash string) (*types.TxResponse, error) {
	clientCtx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return nil, err
	}

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

func QueryTxWithRetry(cmd *cobra.Command, txHash string, trials int) (*types.TxResponse, error) {
	var broadcastedTx *types.TxResponse
	var err error

	for i := 0; i < trials; i++ {
		fmt.Println(i)
		time.Sleep(10 * time.Second)
		broadcastedTx, err = QueryTx(cmd, txHash)
		if err == nil {
			break
		}

	}

	return broadcastedTx, err
}
