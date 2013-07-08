package math

import "testing"

func TestEvenOdd(t *testing.T) {
	for i := 0; i < 50; i++ {
		if IsOdd(2*i) || !IsEven(2*i) {
			t.Errorf("IsEven/IsOdd: test failed with 2*i = %d\n", 2*i)
		}
		if IsEven(2*i + 1) || !IsOdd(2*i + 1) {
			t.Errorf("IsEven/IsOdd: test failed with 2*i + 1 = %d\n", 2*i + 1)
		}
	}
}

func TestGCD(t *testing.T) {
	var d uint64
	if d = Gcd(6, 3); d != 3 {
		t.Errorf("Gcd: test failed: got gcd(6,3) = %q", d)
	}
	if d = Gcd(5, 11); d != 1 {
		t.Errorf("Gcd: test failed: got gcd(5,11) = %q", d)
	}
	if d = Gcd(6, 9); d != 3 {
		t.Errorf("Gcd: test failed: got gcd(3,9) = %q", d)
	}
}

