package invest

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/notional-labs/cookiemonster/internal/invest/phase"
)

type Fund struct {
	Address         sdk.AccAddress
	KeyName         string
	TransferTo      map[string]float32
	PoolPercentage  float32
	StakePercentage float32
	PoolPtrategy    phase.PoolStrategy
}

type Funds []Fund

func (fund Fund) Invest() {
	phase.ClaimRewardFromFund(fund)

}
