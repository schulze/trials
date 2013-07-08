package crypto

import (
	"testing"
	"math/big"
)

func TestProduct(t *testing.T) {
	x := new(big.Int)
	ProductInt64(x, []int64{314,159, 265, 359, 897})
	res, _ := new(big.Int).SetString("4260489878970", 10)
	if !Equal(x, res) {
		t.Fatal(x, res)
	}
}

func TestRho(t *testing.T) {
	// second rho example is faster because prime is smaller:
	x, _ := new(big.Int).SetString("314159265358979323", 10)
	a, b := Rho(x)
	if !Equal(a, big.NewInt(990371647)) || !Equal(b, big.NewInt(317213509)) {
		t.Fatal()
	}

	y, _ := new(big.Int).SetString("698599699288686665490308069057420138223871", 10)
	a, b = Rho(y)
	res1, _ := new(big.Int).SetString("2053", 10)
	res2, _ := new(big.Int).SetString("340282366920938463463374607431768211507", 10)
	if !Equal(a, res1) || !Equal(b, res2) {
		t.Fatal()
	}
}
