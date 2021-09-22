protected computeAPYForSpecificDuration(
    poolId: string,
    duration: Duration,
    priceStore: {
    getPrice(coinId: string, vsCurrency: string): number | undefined;
    calculatePrice(coin: CoinPretty, vsCurrrency?: string): PricePretty | undefined;
},
fiatCurrency: FiatCurrency
): IntPretty {
    const gaugeId = this.getIncentivizedGaugeId(poolId, duration);

    if (this.incentivizedPools.includes(poolId) && gaugeId) {
        /*
         XXX: 현재로서는 이 메소드는 Incentivized Pools 카드에서 사용된다.
              근데 Incentivized Pools 카드는 All Pools 카드와 같은 페이지에서 보이기 때문에
              쿼리 수를 줄이기 위해서 pagination으로부터 풀의 정보를 받아오도록 한다.
         */
        const pool = this.queryPools.getPoolFromPagination(poolId);
        if (pool) {
            const mintDenom = this.queryMintParmas.mintDenom;
            const epochIdentifier = this.queryMintParmas.epochIdentifier;

            if (mintDenom && epochIdentifier) {
                const epoch = this.queryEpochs.getEpoch(epochIdentifier);

                const chainInfo = this.chainGetter.getChain(this.chainId);
                const mintCurrency = chainInfo.findCurrency(mintDenom);

                if (mintCurrency && mintCurrency.coinGeckoId && epoch.duration) {
                    const totalWeight = this.queryDistrInfo.totalWeight;
                    const potWeight = this.queryDistrInfo.getWeight(gaugeId);
                    const mintPrice = priceStore.getPrice(mintCurrency.coinGeckoId, fiatCurrency.currency);
                    const poolTVL = pool.computeTotalValueLocked(priceStore, fiatCurrency);
                    if (totalWeight.gt(new Int(0)) && potWeight.gt(new Int(0)) && mintPrice && poolTVL.toDec().gt(new Dec(0))) {
                        // 에포치마다 발행되는 민팅 코인의 수.
                        const epochProvision = this.queryEpochProvision.epochProvisions;

                        if (epochProvision) {
                            const numEpochPerYear =
                                dayjs
                                    .duration({
                                        years: 1,
                                    })
                                    .asMilliseconds() / epoch.duration.asMilliseconds();

                            const yearProvision = epochProvision.mul(new Dec(numEpochPerYear.toString()));
                            const yearProvisionToPots = yearProvision.mul(
                                this.queryMintParmas.distributionProportions.poolIncentives
                            );
                            const yearProvisionToPot = yearProvisionToPots.mul(new Dec(potWeight).quo(new Dec(totalWeight)));

                            const yearProvisionToPotPrice = new Dec(mintPrice.toString()).mul(yearProvisionToPot.toDec());

                            // 백분률로 반환한다.
                            return new IntPretty(yearProvisionToPotPrice.quo(poolTVL.toDec()))
                                .decreasePrecision(2)
                                .maxDecimals(2)
                                .trim(true);
                        }
                    }
                }
            }
        }
    }
