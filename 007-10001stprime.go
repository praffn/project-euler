// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
//
// What is the 10 001st prime number?
package main

import (
	"fmt"

	"./primes"
)

func main() {
	p := []int{}
	primeCh := make(chan int)
	go primes.Sieve(primeCh, 10001)
	for i := range primeCh {
		p = append(p, i)
	}
	fmt.Println(p[len(p)-1])
}
