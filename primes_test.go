package primes

import (
	"testing"
)

func TestPrimeWheel(t *testing.T) {
	current := int64(2)
	wheel := PrimeWheel()
	for current < int64(1000000000) {
		if current += wheel(); current&1 == 0 {
			t.Errorf("The wheel reached an even number, %d\n", current)
		}
	}
}

func TestIsPrime(t *testing.T) {
	not_primes := [...]int64{10, 33, 121, 2441838, 35100099}
	primes := [...]int64{11, 97, 50647, 104729}
	for _, v := range not_primes {
		if IsPrime(v) {
			t.Errorf("Incorrectly reported %d as prime", v)
		}
	}
	for _, v := range primes {
		if !IsPrime(v) {
			t.Errorf("Incorrectly reported %d as composite", v)
		}
	}
}

func TestPrimeFactorize(t *testing.T) {
	var factors []int64
	var product int64
	for i := int64(2); i <= int64(10000); i++ {
		factors = PrimeFactorize(i)
		product = 1
		for _, n := range factors {
			product *= n
		}
		if product != i {
			t.Errorf("Product of prime factorization of %d equals %d",
				i, product)
		}
	}
}

func TestPrimeSieve(t *testing.T) {
	primes := PrimeSieve(1000000)
	for _, v := range primes {
		if !IsPrime(v) {
			t.Errorf("Prime sieve generated %d, which is not prime", v)
		}
	}
}

func TestLazyPrimes(t *testing.T) {
	var p int64
	primes := make([]int64, 0)
	c := make(chan int64)
	go LazyPrimes(1000000, c)
	for {
		p = <-c
		if p == 0 {
			break
		}
		primes = append(primes, p)
	}
	primes2 := PrimeSieve(1000000)
	for i, v := range primes2 {
		if v != primes[i] {
			t.Errorf("The %dth prime should be %d, got %d", i+1, v, primes[i])
		}
	}
}
