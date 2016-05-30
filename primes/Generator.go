package primes

func Generator() func() int {
	previous := 1
	return func() int {
		for {
			previous++
			if IsPrime(previous) {
				break
			}
		}
		return previous
	}
}
