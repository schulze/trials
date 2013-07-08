package crypto

import (
	"math/big"
)

func IsOne(x *big.Int) bool {
	return Equal(x, big1)
}

func Equal(x, y *big.Int) bool {
	return x.Cmp(y) == 0
}