package main

import (
	"fmt"
	"github.com/pelmers/primes"
	"os"
	"strconv"
)

func main() {
	var limit int64
	if len(os.Args) <= 1 {
		limit = 10000000
	} else {
		limit, _ = strconv.ParseInt(os.Args[1], 10, 64)
	}
	fmt.Println("Computing the biggest prime less than", limit, "...")
	p := primes.PrimeSieve(limit)
	fmt.Println("There are", len(p), "primes less than", limit)
	fmt.Println(p[len(p)-1], "is the biggest prime less than", limit)
}
