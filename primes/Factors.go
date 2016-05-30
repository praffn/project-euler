package primes

import "fmt"

func Factors(num int) (factors []int) {
	primes := make(chan int)
	go Primes(num/2, primes)

	for prime := range primes {
		fmt.Println(prime)
	}

	// prime := Generator()
	// max := num / 2
	//
	// for p := prime(); p <= max; p = prime() {
	// 	fmt.Println(p)
	// 	if num%p == 0 {
	// 		factors = append(factors, p)
	// 	}
	// }

	return factors
}
