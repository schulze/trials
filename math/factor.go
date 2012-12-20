package math

var (
	k = len(SmallPrimes)
	lastPrime = SmallPrimes[k-1]
)

func setIndex() (j int64) {
	pk := lastPrime % 30
	switch pk {
	case 1: return 0
	case 7: return 1
	case 11: return 2
	case 13: return 3
	case 17: return 4
	case 19: return 5
	case 23: return 6
	case 29: return 7
	}
	// can't happen
	return -1
}

// TrialDivide tries do divide n by integers less than b.  If a non-trivial
// divisor is found TrialDivide returns this divisor. Otherwise returns 1.
func TrialDivide(n, b int64) int64 {
	t := []int64{6, 4, 2, 4, 2, 4, 6, 2}

	// divide by the stored small primes first
	for _, p := range SmallPrimes {
		if p >= b {
			return 1
		}
		if n % p == 0 {
			return p
		}
	}

	// divide by numbers bigger than our small prime numbers
	i := setIndex()
	for d := lastPrime; d < b; d += t[i] {
		i += 1
		if n % d == 0 {
			return d
		}
	}
	return 1
}

// IsPrime tests if n is a prime integer.
func IsPrime(n int64) bool {
	d := TrialDivide(n, n)
	if d == 1 {
		return true
	}
	return false
}

