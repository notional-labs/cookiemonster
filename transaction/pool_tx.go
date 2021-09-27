package transaction

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/cookiemonster/osmosis"
	"github.com/osmosis-labs/osmosis/x/gamm/types"
	"gopkg.in/yaml.v3"
)

type JoinPoolOption struct {
	PoolId         uint64
	ShareOutAmount sdk.Int
	MaxAmountsIn   sdk.Coins
}

func NewMsgJoinPool(fromAddr sdk.AccAddress, poolId uint64, shareOutAmount sdk.Int, maxAmountsIn sdk.Coins) sdk.Msg {
	msg := &types.MsgJoinPool{
		Sender:         fromAddr.String(),
		PoolId:         poolId,
		ShareOutAmount: shareOutAmount,
		TokenInMaxs:    maxAmountsIn,
	}
	return msg
}

func JoinPool(keyName string, joinPoolOpt JoinPoolOption) error {
	// build tx context
	clientCtx := osmosis.DefaultClientCtx
	clientCtx, err := SetKeyNameToContext(clientCtx, keyName)
	if err != nil {
		return err
	}
	txf := NewTxFactoryFromClientCtx(clientCtx)

	// build msg for tx
	fromAddr := clientCtx.FromAddress
	poolId := joinPoolOpt.PoolId
	shareOutAmount := joinPoolOpt.ShareOutAmount
	maxAmountsIn := joinPoolOpt.MaxAmountsIn

	msg := NewMsgJoinPool(fromAddr, poolId, shareOutAmount, maxAmountsIn)

	return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
}

type JoinPoolTx struct {
	KeyName     string
	JoinPoolOpt JoinPoolOption
}

func (joinPoolTx JoinPoolTx) Execute() error {

	keyName := joinPoolTx.KeyName
	joinPoolOpt := joinPoolTx.JoinPoolOpt
	err := JoinPool(keyName, joinPoolOpt)
	return err
}

func (joinPoolTx JoinPoolTx) Report() {

	joinPoolOpt := joinPoolTx.JoinPoolOpt
	keyName := joinPoolTx.KeyName

	f, _ := os.OpenFile("report", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	f.WriteString("\nJoin Pool Transaction\n")
	f.WriteString("\nKeyname: " + keyName + "\n")
	f.WriteString("\nJoin Pool Option\n\n")

	txData, _ := yaml.Marshal(joinPoolOpt)
	_, _ = f.Write(txData)
	f.WriteString(transactionSeperator)

	f.Close()
}
