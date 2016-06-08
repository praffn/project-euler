package primes

import "math"

func IsPrime(num int) bool {
	length := int(math.Sqrt(float64(num)))
	for i := 2; i <= length; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
