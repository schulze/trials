// Some factorisations base on FactHacks talk by Daniel J. Bernstein, Nadia Heninger, and Tanja Lange.
package crypto

import (
	"math/big"
)

var (
	big0 = big.NewInt(0)
	big1 = big.NewInt(1)
)

// mulPair sets z to the product of the first two elements in s.
func mulPair(z *big.Int, s []*big.Int) *big.Int {
	if len(s) == 1 {
		return z.Set(s[0])
	}
	return z.Mul(s[0], s[1])
}

// Product sets z to the product of the integers in s and returns z.
func Product(z *big.Int, s []*big.Int) *big.Int {
	if len(s) == 0 {
		return z.SetInt64(1)
	}
	// make sure the length of s is even
	if len(s) % 2 == 1 {
		s[0].Mul(s[0], s[len(s)-1])
		s = s[:len(s)-1]
	}
	for len(s) > 1 {
		p := make([]*big.Int, (len(s)+1)/2)
		for i := 0; i < (len(s)+1)/2; i++ {
			p[i] = &big.Int{}
			mulPair(p[i], s[2*i: 2*(i+1)])
		}
		s = p
	}
	return z.Set(s[0])
}

// ProductInt64 sets z to the product of the integers in s and returns z.
func ProductInt64(z *big.Int, s []int64) *big.Int {
	if len(s) == 0 {
		return z.SetInt64(1)
	}
	bs := make([]*big.Int, len(s))
	for i, v := range s {
		bs[i] = big.NewInt(v)
	}
	return Product(z, bs)
}

// 
func Rho(n *big.Int) (*big.Int, *big.Int) {
	c := big.NewInt(10)
	a0 := big.NewInt(1)
	a1 := new(big.Int)
	a1 = a1.Mul(a0, a0).Add(a1, c)
	a2 := new(big.Int)
	a2 = a2.Mul(a1, a1).Add(a2, c)
	
	d := new(big.Int)
	g := new(big.Int)
	for Equal(g.GCD(nil, nil, n, d.Sub(a2, a1).Abs(d)), big1) {
		a1.Mul(a1, a1).Add(a1, c).Mod(a1, n)
		a2.Mul(a2, a2).Add(a2, c).Mod(a2, n)
		a2.Mul(a2, a2).Add(a2, c).Mod(a2, n)
	}
	return g, n.Div(n, g)
}
