package main

import (
	"context"

	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/notional-labs/cookiemonster/internal/osmosis"
)

func QueryBalances(keyName string) (sdk.Coins, error) {
	// build context
	clientCtx := osmosis.DefaultClientCtx

	queryClient := types.NewQueryClient(clientCtx)
	addr := clientCtx.FromAddress
	pageReq := &query.PageRequest{
		Key:        []byte(""),
		Offset:     uint64(0),
		Limit:      uint64(100),
		CountTotal: false,
	}
	params := types.NewQueryAllBalancesRequest(addr, pageReq)
	res, err := queryClient.AllBalances(context.Background(), params)
	if err != nil {
		return nil, err
	}
	return res.Balances, nil
}

func QueryOsmoBalance(keyName string) (*big.Int, error) {
	// build context
	clientCtx := osmosis.DefaultClientCtx

	queryClient := types.NewQueryClient(clientCtx)
	addr := clientCtx.FromAddress
	params := types.NewQueryBalanceRequest(addr, "uosmo")
	res, err := queryClient.Balance(context.Background(), params)
	if err != nil {
		return nil, err
	}
	osmosBalance := res.Balance.Amount.BigInt()
	return osmosBalance, nil
}
