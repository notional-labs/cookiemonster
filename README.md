# Cookiemonster
Cookiemonster is an open source tool for the automatic managment of liquidity provision and staking in Osmosis, and Soon(TM), other Cosmos IBC enabled blockchains.  Planned feature expansion includes:

- [x] Pool Zero Automation for Osmosis
- [ ] Automation of many pools for Osmosis
- [ ] Validator Management from a disconnected signer
- [ ] Osmo <-> Sif arbitrage
- [ ] Osmo <-> Gdex arbitrage
- [ ] Staking claims at optimal intervals


## Compile it
The ideal way to run Cookie Monster is on a locked-down raspberry Pi that has access to ONLY the API endpoints that cookiemonster needs to consume.  You can use [SOS](https://github.com/notional-labs/sos) and configure its security policies to your liking.  


For the time being, Cookie Monster has hot keys in the filesystem, like a relayer, so secure the machine running Cookie Monster as though it **is** your private keys.  Once agian, note that we've explicitly chosen not to secure Cookie Monster, so that he could be built quickly.  There are ways to secure cookie monster, and they're left to the users, for now. 

```
git clone https://github.com/notional-labs/cookiemonster
cd cookiemonster
git checkout pool_testing
cd cmd/auto-farm
go build
mv auto-farm /usr/local/bin
```


One-liner coming Soon(TM)


## Run it
1. Put your account into auto-farm keyring by : `auto-farm keys add {your account_name} --recover`
2. Create an an accountmanager.json file at ~/.cookiemonster/investments.json, as seen below
```
[
{
    "KeyName":"{your account_name}",
    "TransferTo":{},
    "PoolPercentage":5,
    "StakePercentage":0,
    "PoolStrategy":{"Name":"custom","Config":{"1":100},"ConfigDenom":"percentages"},
    "Duration":"14days",
    "StakeAddress":""
}
]
```
3. run: `auto-invest invest ~/.cookiemonster/investments.json`.  This cmd will run ONE investment : Claim -> Pooling -> Staking. It finish when all transaction is success or 4 times broadcast. </br>
  Or run: `auto-farm invest_non_stop ~/.cookiemonster/investments.json`.   This cmd will run loop of investment for each epoch in osmosis.



## License

MIT
