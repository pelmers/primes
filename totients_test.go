package primes

import "testing"

func TestTotient(t *testing.T) {
	testcases := [...]int64{5, 4, 390, 96, 1930, 768, 13798, 6898, 319808, 150912,
		38170805, 30536640}
	for i := 0; i < len(testcases); i += 2 {
		if Totient(testcases[i]) != testcases[i+1] {
			t.Errorf("Totient of %d should be %d, got %d",
				testcases[i], testcases[i+1], Totient(testcases[i]))
		}
	}
}

func TestTotientSieve(t *testing.T) {
	totes := TotientSieve(10000)
	for i := int64(2); i < int64(len(totes)); i++ {
		if totes[i] != Totient(i) {
			t.Errorf("Totient of %d should be %d, not %d",
				i, Totient(i), totes[i])
		}
	}
}
