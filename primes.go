// euler/primes
//  Provides some tools for generating and checking primes and totients
//
// primes.go:
//  -- PrimeWheel creates a prime wheel that skips multiples of 2,3,5,7
//  -- IsPrime tests the primacy of a number
//  -- PrimeFactorize gives the prime factorization of a number
//  -- UniquePrimeFactors gives the unique prime factors of a number
//  -- PrimeSieve returns a list of all the primes up to a number
//  -- LazyPrimes sends primes to an output channel as they are found
//
// totients.go:
//  -- Totient gives the evaluation of Euler's totient function of a number
//  -- TotientSieve finds the totient of every number up to a number
package primes

import "math"

// Return a function that spins a 2,3,5,7 prime wheel when called.
// The function is a closure that keeps track of where on the wheel it is.
// example usage:
//  possible_prime := 2
//  wheel := PrimeWheel()
//  possible_prime += wheel()
func PrimeWheel() func() int64 {
	current := 0
	holes := ([...]int64{1, 2, 2, 4, 2, 4, 2, 4, 6, 2, 6, 4, 2, 4, 6, 6, 2, 6, 4, 2, 6, 4, 6, 8, 4, 2, 4, 2, 4,
		8, 6, 4, 6, 2, 4, 6, 2, 6, 6, 4, 2, 4, 6, 2, 6, 4, 2, 4, 2, 10, 2, 10})
	return func() int64 {
		if current == len(holes) {
			current = 4
		}
		current++
		return holes[current-1]
	}
}

// Return the primacy of a number.
// true if num is prime, else false.
func IsPrime(num int64) bool {
	wheel := PrimeWheel()
	// we only have to test up to the square root of the number
	r := int64(math.Sqrt(float64(num)))
	for p := int64(2); p <= r; p += wheel() {
		if num%p == 0 {
			return false
		}
	}
	return true
}

// Return an int64 slice of the prime factorization of a number >= 2.
// The prime factorization is sorted, and its product is the original number.
func PrimeFactorize(num int64) []int64 {
	var p int64
	factors := make([]int64, 0)
	primechan := make(chan int64)
	go LazyPrimes(num, primechan)
	// number will equal one at the end
	for num != 1 {
		// note: guarantees that smallest prime factors appear first
		for p = <-primechan; num%p == 0; num /= p {
			factors = append(factors, p)
		}
	}
	return factors
}

// Return an int64 slice of the unique prime factors of a number >= 2.
// The factors in the slice are sorted in increasing order.
func UniquePrimeFactors(num int64) []int64 {
	var p int64
	factors := make([]int64, 0)
	primechan := make(chan int64)
	go LazyPrimes(num, primechan)
	// stop when the number equals one
	for num != 1 {
		if p = <-primechan; num%p == 0 {
			// only append this once
			factors = append(factors, p)
			// keep dividing by this factor while it's divisible
			for num /= p; num%p == 0; num /= p {
			}
		}
	}
	return factors
}

// Return an int64 slice of all the prime numbers less than the limit.
// Implements the Sieve of Eratosthenes.
func PrimeSieve(limit int64) []int64 {
	var i int64
	rootlimit := int64(math.Sqrt(float64(limit)))
	wheel := PrimeWheel()
	sieve := make([]bool, limit+1)
	// starting capacity based on the prime number theorem
	primes := make([]int64, 0, limit/int64(math.Log(float64(limit))))
	for i = 2; i <= rootlimit; i += wheel() {
		if sieve[i] == false {
			primes = append(primes, i)
			for j := i * i; j <= limit; j += i {
				sieve[j] = true
			}
		}
	}
	// look at everything from rootlimit to limit, appending primes
	for ; i <= limit; i += wheel() {
		if sieve[i] == false {
			primes = append(primes, i)
		}
	}
	return primes
}

// Lazily evaluate primes up to the limit, sending primes back on out channel.
// Send 0 when the limit is reached.
func LazyPrimes(limit int64, out chan<- int64) {
	var i int64
	rootlimit := int64(math.Sqrt(float64(limit)))
	wheel := PrimeWheel()
	sieve := make([]bool, limit+1)
	for i = 2; i <= rootlimit; i += wheel() {
		if sieve[i] == false {
			out <- i
			for j := i * i; j <= limit; j += i {
				sieve[j] = true
			}
		}
	}
	for ; i <= limit; i += wheel() {
		if sieve[i] == false {
			out <- i
		}
	}
	// signal end of stream
	out <- 0
}
