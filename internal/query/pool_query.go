package main

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"time"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/notional-labs/cookiemonster/internal/osmosis"
	epoch "github.com/osmosis-labs/osmosis/x/epochs/types"
	"github.com/osmosis-labs/osmosis/x/gamm/types"
)

func QuerySpotPrice(poolId int, tokenInDenom string, tokenOutDenom string) (float64, error) {
	clientCtx := osmosis.DefaultClientCtx

	queryClient := types.NewQueryClient(clientCtx)

	res, err := queryClient.SpotPrice(context.Background(), &types.QuerySpotPriceRequest{
		PoolId:        uint64(poolId),
		TokenInDenom:  tokenInDenom,
		TokenOutDenom: tokenOutDenom,
	})
	if err != nil {
		return 0, err
	}

	spotPriceString := res.GetSpotPrice()
	spotPrice, _ := strconv.ParseFloat(spotPriceString, 64)
	return spotPrice, nil
}

func QueryPools() (*types.QueryPoolsResponse, error) {
	clientCtx := osmosis.DefaultClientCtx

	queryClient := types.NewQueryClient(clientCtx)

	pageReq := &query.PageRequest{
		Key:        []byte(""),
		Offset:     uint64(0),
		Limit:      uint64(100),
		CountTotal: false,
	}

	res, err := queryClient.Pools(context.Background(), &types.QueryPoolsRequest{
		Pagination: pageReq,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func QueryPoolId(poolId int) (*types.QueryPoolResponse, error) {
	clientCtx := osmosis.DefaultClientCtx
	queryClient := types.NewQueryClient(clientCtx)

	res, err := queryClient.Pool(context.Background(), &types.QueryPoolRequest{
		PoolId: uint64(poolId),
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func QueryEpochProvision(epoch epoch.EpochInfo) cosmostypes.Dec {
	//var epochProvision int64 = int64(821917808219178082191780821917) //from API: 821917808219.178082191780821917
	return cosmostypes.MustNewDecFromStr("821917808219178082191780821917")
}

/**
This is just test data.
*/
func getEpoch(epochIdentifier string) epoch.EpochInfo {
	return epoch.EpochInfo{
		Identifier:            "OK",
		StartTime:             time.Now(),
		Duration:              time.Duration(1234123413),
		CurrentEpoch:          int64(1234),
		CurrentEpochStartTime: time.Now(),
		EpochCountingStarted:  true,
	}
}

type DistributionProportions struct {
	PoolIncentives cosmostypes.Dec
}

type Pool struct {
	TotalValueLocked        cosmostypes.Dec // probably wrong.
	DistributionProportions DistributionProportions
	GuageId                 int32
	PotWeight               cosmostypes.Dec
	PoolId                  int32
	PoolIncentives          cosmostypes.Dec
	EpochIdentifier         string
	TotalWeight             cosmostypes.Dec // total weight of funds in poolId
	Duration                int64           // duration in seconds that funds are locked in pool.
	APY                     float64
}

var samplePoolIncentive = 0.45

//Need different pool for each duration or to change this model. I think it makes sense for every time to have it's own pool.
var pool = Pool{
	TotalValueLocked: cosmostypes.NewDec(1004366),
	DistributionProportions: DistributionProportions{
		PoolIncentives: cosmostypes.NewDec(1234),
	},
	PoolId:          1,
	GuageId:         1,
	PotWeight:       cosmostypes.NewDec(359034), //359034n
	TotalWeight:     cosmostypes.NewDec(1004366),
	Duration:        86400,
	EpochIdentifier: "day",
	PoolIncentives:  cosmostypes.NewDec(int64(float64(samplePoolIncentive) * float64(math.Pow(10, 18)))),
	APY:             0,
}

/**
 osmosisd q mint params --node http://95.217.196.54:2001

distribution_proportions:
  community_pool: "0.050000000000000000"
  developer_rewards: "0.250000000000000000"
  pool_incentives: "0.450000000000000000"
  staking: "0.250000000000000000"
epoch_identifier: day
genesis_epoch_provisions: "821917808219.178082191780821917"
mint_denom: uosmo
minting_rewards_distribution_start_epoch: "1"
reduction_factor: "0.666666666666666666"
reduction_period_in_epochs: "365"

*/

func toFloat64(d cosmostypes.Dec) float64 {
	if value, err := strconv.ParseFloat(d.String(), 64); err != nil {
		panic(err)
	} else {
		return value
	}
}

func CalculatePoolAPY(pool Pool, duration time.Duration) Pool {

	// From API /osmosis/pool-incentives/v1beta1/incentivized_pools
	// From API `/osmosis/pool-incentives/v1beta1/distr_info
	var totalWeight = pool.TotalWeight
	//const gaugeId = this.getIncentivizedGaugeId(poolId, duration);
	var potWeight = pool.PotWeight //const potWeight = this.queryDistrInfo.getWeight(gaugeId); comes from distr_info.
	//var poolIncentives = pool.PoolIncentives; //See above, we need to get mint params.
	var oneYearMilliseconds int64 = 365 * 24 * 60 * 60 * 1000
	var epochIdentifier = "OK"
	var epoch = getEpoch(epochIdentifier) // still not sure how to get a valid epochID
	// example epochProvision response: 821917808219.178082191780821917
	var epochProvision = QueryEpochProvision(epoch) //osmosisd q mint epoch-provisions --node http://95.217.196.54:2001
	fmt.Println("EpochProvision")
	fmt.Println(epochProvision)
	var numEpochPerYear = oneYearMilliseconds / epoch.Duration.Milliseconds()
	var yearProvision cosmostypes.Dec = epochProvision.Mul(cosmostypes.NewDec(numEpochPerYear))

	//var yearProvisionToPots = yearProvision.Mul(poolIncentives)
	var yearProvisionToPot = yearProvision.Mul(potWeight.Quo(totalWeight))
	var poolTLV = pool.TotalValueLocked // should be 821917808219178082191780821917n
	fmt.Println("PoolTLV")
	fmt.Println(poolTLV)
	//return new IntPretty(yearProvisionToPotPrice.quo(poolTVL.toDec()))
	//.decreasePrecision(2)
	//.maxDecimals(2)
	//.trim(true);
	var APY = yearProvisionToPot.Quo(pool.TotalValueLocked)

	pool.APY = toFloat64(APY)
	fmt.Println(pool.APY)
	return pool
}

func main() {
	CalculatePoolAPY(pool, time.Duration(100))
}

//const poolTVL = pool.computeTotalValueLocked(priceStore, fiatCurrency);
//if (totalWeight.gt(new Int(0)) && potWeight.gt(new Int(0)) && mintPrice && poolTVL.toDec().gt(new Dec(0))) {
//// 에포치마다 발행되는 민팅 코인의 수.
// API: osmosis/mint/v1beta1/epoch_provisions
//const epochProvision = this.queryEpochProvision.epochProvisions;
//
//if (epochProvision) {
//const numEpochPerYear =
//dayjs
//.duration({
//years: 1,
//})
//.asMilliseconds() / epoch.duration.asMilliseconds();
//

//const totalWeight = this.queryDistrInfo.totalWeight;
//						const potWeight = this.queryDistrInfo.getWeight(gaugeId);
//const yearProvision = epochProvision.mul(new Dec(numEpochPerYear.toString()));
//const yearProvisionToPots = yearProvision.mul(
//this.queryMintParmas.distributionProportions.poolIncentives
//);
//const yearProvisionToPot = yearProvisionToPots.mul(new Dec(potWeight).quo(new Dec(totalWeight)));
//
//const yearProvisionToPotPrice = new Dec(mintPrice.toString()).mul(yearProvisionToPot.toDec());
//
//// 백분률로 반환한다.
//return new IntPretty(yearProvisionToPotPrice.quo(poolTVL.toDec()))
//.decreasePrecision(2)
//.maxDecimals(2)
//.trim(true);
//}

/**
incentivized_pools:
- gauge_id: "1"
  lockable_duration: 86400s
  pool_id: "1"
- gauge_id: "2"
  lockable_duration: 604800s
  pool_id: "1"
- gauge_id: "3"
  lockable_duration: 1209600s
  pool_id: "1"
- gauge_id: "4"
  lockable_duration: 86400s
  pool_id: "2"
- gauge_id: "5"
  lockable_duration: 604800s
  pool_id: "2"
- gauge_id: "6"
  lockable_duration: 1209600s
  pool_id: "2"
- gauge_id: "7"
  lockable_duration: 86400s
  pool_id: "3"
- gauge_id: "8"
  lockable_duration: 604800s
  pool_id: "3"
- gauge_id: "9"
  lockable_duration: 1209600s
  pool_id: "3"
- gauge_id: "10"
  lockable_duration: 86400s
  pool_id: "4"
- gauge_id: "11"
  lockable_duration: 604800s
  pool_id: "4"
- gauge_id: "12"
  lockable_duration: 1209600s
  pool_id: "4"
- gauge_id: "13"
  lockable_duration: 86400s
  pool_id: "5"
- gauge_id: "14"
  lockable_duration: 604800s
  pool_id: "5"
- gauge_id: "15"
  lockable_duration: 1209600s
  pool_id: "5"
- gauge_id: "16"
  lockable_duration: 86400s
  pool_id: "6"
- gauge_id: "17"
  lockable_duration: 604800s
  pool_id: "6"
- gauge_id: "18"
  lockable_duration: 1209600s
  pool_id: "6"
- gauge_id: "19"
  lockable_duration: 86400s
  pool_id: "7"
- gauge_id: "20"
  lockable_duration: 604800s
  pool_id: "7"
- gauge_id: "21"
  lockable_duration: 1209600s
  pool_id: "7"
- gauge_id: "22"
  lockable_duration: 86400s
  pool_id: "8"
- gauge_id: "23"
  lockable_duration: 604800s
  pool_id: "8"
- gauge_id: "24"
  lockable_duration: 1209600s
  pool_id: "8"
- gauge_id: "25"
  lockable_duration: 86400s
  pool_id: "9"
- gauge_id: "26"
  lockable_duration: 604800s
  pool_id: "9"
- gauge_id: "27"
  lockable_duration: 1209600s
  pool_id: "9"
- gauge_id: "28"
  lockable_duration: 86400s
  pool_id: "10"
- gauge_id: "29"
  lockable_duration: 604800s
  pool_id: "10"
- gauge_id: "30"
  lockable_duration: 1209600s
  pool_id: "10"
- gauge_id: "37"
  lockable_duration: 86400s
  pool_id: "13"
- gauge_id: "38"
  lockable_duration: 604800s
  pool_id: "13"
- gauge_id: "39"
  lockable_duration: 1209600s
  pool_id: "13"
- gauge_id: "43"
  lockable_duration: 86400s
  pool_id: "15"
- gauge_id: "44"
  lockable_duration: 604800s
  pool_id: "15"
- gauge_id: "45"
  lockable_duration: 1209600s
  pool_id: "15"
- gauge_id: "64"
  lockable_duration: 86400s
  pool_id: "22"
- gauge_id: "65"
  lockable_duration: 604800s
  pool_id: "22"
- gauge_id: "66"
  lockable_duration: 1209600s
  pool_id: "22"
- gauge_id: "124"
  lockable_duration: 86400s
  pool_id: "42"
- gauge_id: "125"
  lockable_duration: 604800s
  pool_id: "42"
- gauge_id: "126"
  lockable_duration: 1209600s
  pool_id: "42"
- gauge_id: "558"
  lockable_duration: 86400s
  pool_id: "183"
- gauge_id: "559"
  lockable_duration: 604800s
  pool_id: "183"
- gauge_id: "560"
  lockable_duration: 1209600s
  pool_id: "183"
- gauge_id: "600"
  lockable_duration: 86400s
  pool_id: "197"
- gauge_id: "601"
  lockable_duration: 604800s
  pool_id: "197"
- gauge_id: "602"
  lockable_duration: 1209600s
  pool_id: "197"
*/
