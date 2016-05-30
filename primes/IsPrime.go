package primes

import "math"

func IsPrime(value int) bool {
	x := float64(value)

	if x <= 1 {
		return false
	}

	if math.Mod(x, 2) == 0 {
		return x == 2
	}

	for i := 3.0; i <= math.Floor(math.Sqrt(x)); i += 2.0 {
		if math.Mod(x, i) == 0 {
			return false
		}
	}

	return true
}
