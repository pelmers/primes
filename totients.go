package primes

// Return the evaluation of Euler's totient function of a number.
// The totient of a number is the number of numbers (< number) coprime to it
func Totient(num int64) int64 {
	factors := UniquePrimeFactors(num)
	for _, v := range factors {
		num = (num / v) * (v - 1)
	}
	return num
}

func TotientSieve(upper int64) []int64 {
	wheel := PrimeWheel()
	totes := make([]int64, upper)
	for i := int64(0); i < int64(len(totes)); i++ {
		totes[i] = i
	}
	for i := int64(2); i < upper; i += wheel() {
		if totes[i] == i {
			for j := i; j < upper; j += i {
				totes[j] = (totes[j] / i) * (i - 1)
			}
		}
	}
	return totes
}
