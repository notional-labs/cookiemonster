


## Compile it


```
git clone https://github.com/notional-labs/cookiemonster
cd cookiemonster
git checkout pool_testing
cd cmd/auto-farm
go install .
```


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
```
KeyName is the you use to sign tx
PoolPercentage is the percentage of your fund you use for pooling, in short it dictates pooling amount
StakePercentage is similar but for staking amount
PoolStrategy is how you pool, namely "Config" field is a map of pool id -> percentage of your pooling amount that you use to pool in that pool id
Duration is the bonding period u choose
StakeAddresss is valoper address you delegate to
```

3. run: `auto-invest invest ~/.cookiemonster/investments.json`.  This cmd will run ONE investment : Claim -> Pooling -> Staking. It finish when all transaction is success or 4 times broadcast. </br>
  Or run: `auto-farm invest_non_stop ~/.cookiemonster/investments.json`.   This cmd will run loop of investment for each epoch in osmosis.



## License

MIT
