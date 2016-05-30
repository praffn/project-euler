package primes

func square(n int) int {
	return n * n
}

func LargestFactor(val int, divisor int) int {
	for val%divisor != 0 && square(divisor) <= val {
		divisor++
	}
	if square(divisor) <= val {
		return LargestFactor(val/divisor, divisor)
	}
	return val
}
