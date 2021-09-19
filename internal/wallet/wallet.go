package wallet

type DispursementStrategy struct {
	name string
}

type RewardSourceType int

const (
	LiquidityPool RewardSourceType = iota
	Staking
)

func (s RewardSourceType) String() string {
	switch s {
	case LiquidityPool:
		return "LiquidityPool"
	case Staking:
		return "Staking"
	}
	return "unknown"
}

type RewardSource struct {
	name       string
	network    string
	rewardtype RewardSourceType
	identifier string
}

type Wallet struct {
	network              string
	address              string
	dispursementStrategy DispursementStrategy
	key                  string
	rewardSources        []RewardSource
}
