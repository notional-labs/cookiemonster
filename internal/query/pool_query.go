package query

import (
	"github.com/cosmos/cosmos-sdk/client"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	epoch "github.com/osmosis-labs/osmosis/x/epochs/types"
	gammcli "github.com/osmosis-labs/osmosis/x/gamm/client/cli"
	"github.com/osmosis-labs/osmosis/x/gamm/types"
	"strconv"
	"time"
)

func QuerySpotPrice(poolId int, tokenInDenom string, tokenOutDenom string) (float64, error) {
	cmd := gammcli.GetCmdSpotPrice()
	clientCtx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return 0, err
	}
	queryClient := types.NewQueryClient(clientCtx)

	if err != nil {
		return 0, err
	}

	res, err := queryClient.SpotPrice(cmd.Context(), &types.QuerySpotPriceRequest{
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

func QueryEpochProvision(epoch epoch.EpochInfo) cosmostypes.Dec {
	return cosmostypes.NewDec(1234)
}



func getEpoch(epochIdentifier string) epoch.EpochInfo {
	return epoch.EpochInfo{
		Identifier: "OK",
		StartTime: time.Now(),
		Duration: time.Duration(1234123413),
		CurrentEpoch: int64(1234),
		CurrentEpochStartTime: time.Now(),
		EpochCountingStarted: true,
	}
}

type DistributionProportions struct {
	PoolIncentives cosmostypes.Dec
}

type Pool struct {
	TotalValueLocked cosmostypes.Dec // probably wrong.
	DistributionProportions DistributionProportions
}

func QueryAPY(pool Pool, duration time.Duration) cosmostypes.Dec {


	// From API `/osmosis/pool-incentives/v1beta1/distr_info
	var totalWeight = cosmostypes.NewDec(1.0); // Weight 50/50.  Need to actually get from Pool.
	//const gaugeId = this.getIncentivizedGaugeId(poolId, duration);
	var potWeight = cosmostypes.NewDec(0.5) //const potWeight = this.queryDistrInfo.getWeight(gaugeId);
	var poolIncentives = cosmostypes.NewDec(100); //this.queryMintParmas.distributionProportions.poolIncentives
	var oneYearMilliseconds int64 = 365*24*60*60*1000;
	var epochIdentifier = "OK";
	var epoch = getEpoch(epochIdentifier);
	var epochProvision = QueryEpochProvision(epoch);
	var numEpochPerYear = oneYearMilliseconds / epoch.Duration.Milliseconds();
	var yearProvision cosmostypes.Dec = epochProvision.Mul(cosmostypes.NewDec(numEpochPerYear))

	var yearProvisionToPots = yearProvision.Mul(poolIncentives)
	var yearProvisionToPot = yearProvision.Mul(potWeight.Quo(totalWeight));
	var APY = yearProvisionToPot.Quo(pool.TotalValueLocked)
	return APY // this is a Dec, need to figure out if we should adjust type.
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