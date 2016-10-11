package util

import "math"

func Divisors(n uint) (c uint) {
	for i := uint(1); i <= uint(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			c = c + 2
		}
	}
	return c
}
