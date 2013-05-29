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
