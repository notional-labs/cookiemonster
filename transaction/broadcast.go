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
func BroadcastTx(clientCtx client.Context, txf clienttx.Factory, msgs ...sdk.Msg) (uint32, string, error) {
	txf, err := clienttx.PrepareFactory(clientCtx, txf)

	if err != nil {
		return 3, "", err
	}

	if clientCtx.Simulate {
		return 3, "", nil
	}

	tx, err := clienttx.BuildUnsignedTx(txf, msgs...)
	if err != nil {
		return 3, "", err
	}

	if !clientCtx.SkipConfirm {
		out, err := clientCtx.TxConfig.TxJSONEncoder()(tx.GetTx())
		if err != nil {
			return 3, "", err
		}

		_, _ = fmt.Fprintf(os.Stderr, "%s\n\n", out)

		buf := bufio.NewReader(os.Stdin)
		ok, err := input.GetConfirmation("confirm transaction before signing and broadcasting", buf, os.Stderr)

		if err != nil || !ok {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", "cancelled transaction")
			return 3, "", err
		}
	}

	err = clienttx.Sign(txf, clientCtx.GetFromName(), tx, true)
	if err != nil {
		return 3, "", err
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(tx.GetTx())
	if err != nil {
		return 3, "", err
	}

	// broadcast to a Tendermint node
	res, err := clientCtx.BroadcastTx(txBytes)
	if err != nil {
		return 3, "", err
	}

	return res.Code, res.TxHash, clientCtx.PrintProto(res)
}

// for more info
// https://youtu.be/unRldLdllZ8
