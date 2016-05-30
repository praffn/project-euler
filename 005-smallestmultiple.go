// 2520 is the smallest number that can be divided by each of
// the numbersfrom 1 to 10 without any remainder.
//
// What is the smallest positive number that is evenly divisible
// by all of the numbers from 1 to 20?
package main

import "fmt"

func main() {
	max := 20
	i := max
	var r int

	for {
		divisible := true
		for j := 2; j <= max; j++ {
			if i%j != 0 {
				divisible = false
				break
			}
		}
		if divisible {
			r = i
			break
		}
		i++
	}

	fmt.Println(r)
}
