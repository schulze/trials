package math

import (
	"math"
)

// IsEven returns true if x is an even number.
func IsEven(x int) bool {
	return x%2 == 0
}

// IsOdd returns true if x is an odd number.
func IsOdd(x int) bool {
	return x%2 == 1
}

// IsSquare returns true if x is the square of a natural number.
func IsSquare(x int) bool {
	n := int(math.Sqrt(float64(x)))
	return n*n == x
}

// IntSqrt returns floor(sqrt(x).
func IntSqrt(x int) int {
	n := int(math.Sqrt(float64(x)))
	return n
}

// Abs returns the absolut value of x, if this value can be represented by an int.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Gcd returns the greatest commen divisor of a and b,
func Gcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Kronecker returns the Kronecker symbol (a/b).
// See Cohen's 'A Course in Computational Algebraic Number Theory'.
func Kronecker(a, b int) int {
	var tab = [8]int{0, 1, 0, -1, 0, -1, 0, 1}
	var v, k int

	if b == 0 {
		if Abs(a) != 1 { return 0 }
		return 1
	}

	if IsEven(a) && IsEven(b) {
		return 0
	}

	for v = 0; IsEven(b); v++ {
		b /= 2
	}
	if IsEven(v) {
		k = 1
	} else {
		k = tab[a&7] // k = (-1)^((a²-1)/8)
	}
	if b < 0 {
		b = -b
		if a < 0 { k = -k }
	}
	for {
		if a == 0 {
			if b >= 2 {
				return 0
			}
			return k
		}
		for v = 0; IsEven(a); v++ {
			a /= 2
		}
		if IsOdd(v) {
			k = k*tab[b&7]
		}

		if a&b&2 != 0 { // k = (-1)^((a-1)*(b-1)/4) * k
			k = -k
		}
		r := Abs(a)
		a = b%r
		b = r
	}
	return 0
}

