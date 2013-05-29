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