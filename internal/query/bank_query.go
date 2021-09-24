package query
//
//import (
//	"github.com/cosmos/cosmos-sdk/client"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/cosmos/cosmos-sdk/types/query"
//	"github.com/cosmos/cosmos-sdk/x/bank/types"
//	"github.com/notional-labs/cookiemonster/internal/transaction"
//
//	bankcli "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
//)
//
//func QueryBalances(keyName string) (sdk.Coins, error) {
//	cmd := bankcli.GetBalancesCmd()
//
//	// build context
//	clientCtx := client.GetClientContextFromCmd(cmd)
//	transaction.SetContextFromKeyName(clientCtx, keyName)
//	queryClient := types.NewQueryClient(clientCtx)
//	addr := clientCtx.FromAddress
//	pageReq := &query.PageRequest{
//		Key:        []byte(""),
//		Offset:     uint64(0),
//		Limit:      uint64(100),
//		CountTotal: false,
//	}
//	params := types.NewQueryAllBalancesRequest(addr, pageReq)
//	res, err := queryClient.AllBalances(cmd.Context(), params)
//	if err != nil {
//		return nil, err
//	}
//	return res.Balances, nil
//}
//
//func QueryOsmoBalance(keyName string) (sdk.Coin, error) {
//	cmd := bankcli.GetBalancesCmd()
//
//	// build context
//	clientCtx := client.GetClientContextFromCmd(cmd)
//	transaction.SetContextFromKeyName(clientCtx, keyName)
//	queryClient := types.NewQueryClient(clientCtx)
//	addr := clientCtx.FromAddress
//
//	params := types.NewQueryBalanceRequest(addr, "uosmo")
//	res, err := queryClient.Balance(cmd.Context(), params)
//	if err != nil {
//		return sdk.Coin{}, err
//	}
//
//	return *res.Balance, nil
//}
