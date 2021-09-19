package transaction

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type TxOption struct {
	Node     string
	FromAddr sdk.AccAddress
	KeyName  string
	ChainId  string
}
