package invest

import "math/big"

// Cal x percent of a
func XPercentageOf(a *big.Int, x int) *big.Int {
	out := &big.Int{}
	out.Mul(a, big.NewInt(int64(x)))

	out.Div(out, big.NewInt(100))

	return out
}
