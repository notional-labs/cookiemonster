package transaction

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/input"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BroadcastTx attempts to generate, sign and broadcast a transaction with the
// given set of messages. It will also simulate gas requirements if necessary.
// It will return an error upon failure.
func BroadcastTx(clientCtx client.Context, txf clienttx.Factory, msgs ...sdk.Msg) (uint32, string, string, error) {
	txf, err := clienttx.PrepareFactory(clientCtx, txf)
	if err != nil {
		return 3, "", "", err
	}

	if txf.SimulateAndExecute() || clientCtx.Simulate {
		_, adjusted, err := clienttx.CalculateGas(clientCtx.QueryWithData, txf, msgs...)
		if err != nil {
			return 3, "", "", err
		}

		txf = txf.WithGas(adjusted)
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", clienttx.GasEstimateResponse{GasEstimate: txf.Gas()})
	}

	if clientCtx.Simulate {
		return 3, "", "", nil
	}

	tx, err := clienttx.BuildUnsignedTx(txf, msgs...)
	if err != nil {
		return 3, "", "", err
	}

	if !clientCtx.SkipConfirm {
		out, err := clientCtx.TxConfig.TxJSONEncoder()(tx.GetTx())
		if err != nil {
			return 3, "", "", err
		}

		_, _ = fmt.Fprintf(os.Stderr, "%s\n\n", out)

		buf := bufio.NewReader(os.Stdin)
		ok, err := input.GetConfirmation("confirm transaction before signing and broadcasting", buf, os.Stderr)

		if err != nil || !ok {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", "cancelled transaction")
			return 3, "", "", err
		}
	}

	err = clienttx.Sign(txf, clientCtx.GetFromName(), tx, true)
	if err != nil {
		return 3, "", "", err
	}

	fmt.Printf("gas = %d \n", tx.GetTx().GetGas())
	fmt.Println(tx.GetTx().GetMsgs())

	txBytes, err := clientCtx.TxConfig.TxEncoder()(tx.GetTx())
	if err != nil {
		return 3, "", "", err
	}

	// broadcast to a Tendermint node
	res, err := clientCtx.BroadcastTx(txBytes)
	if err != nil {
		return 3, "", "", err
	}

	return res.Code, res.RawLog, res.TxHash, clientCtx.PrintProto(res)
}

// for more info
// https://youtu.be/unRldLdllZ8
