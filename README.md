# Cookiemonster
Cookiemonster is an open source tool for the automatic managment of liquidity provision and staking in Osmosis, and Soon(TM), other Cosmos IBC enabled blockchains.  Planned feature expansion includes:

- [x] Validator Management from a disconnected signer
- [x] Osmo <-> Sif arbitrage
- [x] Osmo <-> Gdex arbitrage
- [x] Staking claims at optimal intervals


## Compile it
The ideal way to run Cookie Monster is on a locked-down raspberry Pi that has access to ONLY the API endpoints that cookiemonster needs to consume.  You can use [SOS](https://github.com/notional-labs/sos) and configure its security policies to your liking.  


For the time being, Cookie Monster has hot keys in the filesystem, like a relayer, so secure the machine running Cookie Monster as though it **is** your private keys.  Once agian, note that we've explicitly chosen not to secure Cookie Monster, so that he could be built quickly.  There are ways to secure cookie monster, and they're left to the users, for now. 


## Run it
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
