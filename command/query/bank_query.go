package query

import (
	"context"
	"math/big"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/spf13/cobra"
)

func QueryBalances(cmd *cobra.Command, keyName string) (sdk.Coins, error) {
	// build context
	clientCtx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return nil, err
	}

	addr, err := GetAddressFromKey(clientCtx, keyName)
	if err != nil {
		return nil, err
	}

	queryClient := types.NewQueryClient(clientCtx)

	pageReq := &query.PageRequest{
		Key:        []byte(""),
		Offset:     uint64(0),
		Limit:      uint64(100),
		CountTotal: false,
	}
	params := types.NewQueryAllBalancesRequest(addr, pageReq)
	res, err := queryClient.AllBalances(cmd.Context(), params)
	if err != nil {
		return nil, err
	}
	return res.Balances, nil
}

func QueryUosmoBalance(cmd *cobra.Command, keyName string) (*big.Int, error) {
	// build context
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return nil, err
	}

	addr, err := GetAddressFromKey(clientCtx, keyName)
	if err != nil {
		return nil, err
	}

	queryClient := types.NewQueryClient(clientCtx)
	params := types.NewQueryBalanceRequest(addr, "uosmo")
	res, err := queryClient.Balance(context.Background(), params)
	if err != nil {
		return nil, err
	}
	osmosBalance := res.Balance.Amount.BigInt()
	return osmosBalance, nil
}
