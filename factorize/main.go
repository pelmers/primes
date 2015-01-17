package main

import (
	"fmt"
	"github.com/pelmers/primes"
	"os"
	"strconv"
)

func main() {
	var num int64
	if len(os.Args) <= 1 {
		return
	}
	for i := 1; i < len(os.Args); i++ {
		num, _ = strconv.ParseInt(os.Args[i], 10, 64)
		fmt.Println(num, ":", primes.PrimeFactorize(num))
	}
}
