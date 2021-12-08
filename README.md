# Cookiemonster
Cookiemonster is an open source tool for the automatic managment of liquidity provision and staking in Osmosis, and Soon(TM), other Cosmos IBC enabled blockchains.  Planned feature expansion includes:

* Validator Management from a disconnected signer
* Osmo <-> Sif arbitrage
* Osmo <-> Gdex arbitrage
* Staking claims at optimal intervals




## How to run server
1. Create an an accountmanager.json file at ~/.cookiemonster/accountmanager.json, as seen below


```json
    {
    "MasterKeyHex": "0123456701234567012345670123456701234567012345670123456701234567",
    "NumOfAccount": 0
    }
```

2. run: `auto-farm initapi`



## License

MIT
