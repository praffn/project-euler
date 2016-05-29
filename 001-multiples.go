// If we list all the natural numbers below 10 that are multiples of 3 or 5,
// we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.
package main

import (
	"fmt"

	"./util"
)

func multiplesOf(roof int, bases []int) []int {
	multiples := []int{}
	for i := 1; i < roof; i++ {
		for _, base := range bases {
			if i%base == 0 {
				multiples = append(multiples, i)
				break
			}
		}
	}
	return multiples
}

func main() {
	bases := []int{3, 5}
	multiples := multiplesOf(1000, bases)
	fmt.Println(util.Sum(multiples))
}
