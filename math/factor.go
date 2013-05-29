package math

var (
	k = len(SmallPrimes)
	lastPrime = SmallPrimes[k-1]
	 t = []int64{6, 4, 2, 4, 2, 4, 6, 2}
	j = getIndex(lastPrime)
)

func getIndex(p int64) (j int64) {
	pk := p % 30
	switch pk {
	case 1: j = 0
	case 7: j = 1
	case 11: j = 2
	case 13: j = 3
	case 17: j = 4
	case 19: j = 5
	case 23: j = 6
	case 29: j = 7
	}
	return 
}

// TrialDivide tries do divide n by integers less than b.  If a non-trivial
// divisor is found TrialDivide returns this divisor; otherwise returns 1.
func TrialDivide(n, b int64) int64 {
	// first divide by the stored small primes first
	for _, p := range SmallPrimes {
		if p >= b {
			return 1
		}
		if n % p == 0 {
			return p
		}
	}

	// divide by numbers bigger than our small prime numbers
	d := lastPrime; i := j
	for d < b {
		d = d + t[i]; i = (i+1) % 8
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

func PollardRho(n int) {}
