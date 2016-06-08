// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.
package main

import (
	"fmt"

	"./primes"
)

func main() {
	max := 2000000
	var i, sum int = 3, 2
	for ; i < max; i++ {
		if primes.IsPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
