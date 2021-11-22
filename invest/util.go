package invest

import (
	"math/big"
	"time"

	"github.com/notional-labs/cookiemonster/command/query"
	"github.com/osmosis-labs/osmosis/x/epochs/types"
	"github.com/spf13/cobra"
)

//cal x percent of a
func XPercentageOf(a *big.Int, x int) *big.Int {
	out := &big.Int{}
	out.Mul(a, big.NewInt(int64(x)))
	out.Div(out, big.NewInt(100))
	return out
}

func CalEpochRemainingTime(cmd *cobra.Command) (time.Duration, error) {
	res, err := query.QueryEpoch(cmd)
	if err != nil {
		return 0, err
	}

	var dayEpoch types.EpochInfo
	for _, epoch := range res.GetEpochs() {
		if epoch.Identifier == "day" {
			dayEpoch = epoch
		}
	}

	currentTime := time.Now()

	epochStartTime := dayEpoch.GetCurrentEpochStartTime()

	remainingEpochTime := currentTime.Sub(epochStartTime)

	return remainingEpochTime, nil
}
