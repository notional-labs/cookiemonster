go 1.15

module github.com/notional-labs/auto-farm

require (
	github.com/cosmos/cosmos-sdk v0.44.0
	github.com/gogo/protobuf v1.3.3
	github.com/spf13/pflag v1.0.5
	github.com/tendermint/tendermint v0.34.12
	github.com/osmosis-labs/osmosis-sdk v0.42.10
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
replace github.com/osmosis-labs/osmosis-sdk => github.com/osmosis-labs/cosmos-sdk v0.42.10-0.20210915013958-01114e89a579
