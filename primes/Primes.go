package primes

func Primes(max int, c chan<- int) {
	for i := 1; i <= max; i++ {
		if IsPrime(i) {
			c <- i
		}
	}
	close(c)
}
