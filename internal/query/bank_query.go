package query

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/notional-labs/cookiemonster/internal/transaction"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func QueryOsmoBalance(keyName string) {
	// build tx context
	clientCtx := client.Context{}
	transaction.SetContextFromKeyName(clientCtx, keyName)
	txf := transaction.NewFactoryCLI(clientCtx)

	queryClient := types.NewQueryClient(clientCtx)

	addr := clientCtx.FromAddress
	if err != nil {
		return err
	}

	pageReq, err := client.ReadPageRequest(cmd.Flags())
	if err != nil {
		return err
	}

	if denom == "" {
		params := types.NewQueryAllBalancesRequest(addr, pageReq)

		res, err := queryClient.AllBalances(cmd.Context(), params)
		if err != nil {
			return err
		}
		return clientCtx.PrintProto(res)
	}

	params := types.NewQueryBalanceRequest(addr, denom)
	res, err := queryClient.Balance(cmd.Context(), params)
	if err != nil {
		return err
	}

	return clientCtx.PrintProto(res.Balance)
}

func ReadPageRequest() (*query.PageRequest, error) {
	pageKey := ""
	offset := 0
	limit := 100
	countTotal := false
	page := 1

	if page > 1 && offset > 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "page and offset cannot be used together")
	}

	if page > 1 {
		offset = (page - 1) * limit
	}

	return &query.PageRequest{
		Key:        []byte(pageKey),
		Offset:     uint64(offset),
		Limit:      uint64(limit),
		CountTotal: countTotal,
	}, nil
}
