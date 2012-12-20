package math

import (
	"testing"
)

func TestSmall(t *testing.T) {
	// test with some small primes
	if TrialDivide(17*23*23, 17*23*23+1) != 17 {
		t.Fatal("17*23*23")
	}
	if TrialDivide(17,19) != 17 {
		t.Fatal("17, 19")	
	}
	if TrialDivide(17,18) != 17 {
		t.Fatal("17, 18")	
	}
	if TrialDivide(17,17) != 1 {
		t.Fatal("17,17")	
	}
	if TrialDivide(17,16) != 1 {
		t.Fatal("17,16")	
	}

	// test with some bigger primes
	p1 := SmallPrimes[k-1]
	p2 := SmallPrimes[k-2]
	if TrialDivide(p1*p2, p1 + 2) != p2 {
		t.Fail()
	}
	if TrialDivide(p1*p2, p1 + 1) != p2 {
		t.Fail()
	}
	if TrialDivide(p1*p2, p2) != 1 {
		t.Fail()
	}
	if TrialDivide(p1*p2, p2 - 1) != 1 {
		t.Fail()
	}
	if TrialDivide(p1*5, p1) != 5 {
		t.Fail()
	}
	if TrialDivide(p1*p2*3, p1) != 3 {
		t.Fail()
	}
}