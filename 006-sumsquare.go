// The sum of the squares of the first ten natural numbers is,
//
// 1^2 + 2^2 + ... + 10^2 = 385
//
// The square of the sum of the first ten natural numbers is,
//
// (1 + 2 + ... + 10)^2 = 55^2 = 3025
//
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
//
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
package main

import (
	"fmt"

	"./util"
)

func main() {
	max := 100
	numbers := []int{}
	for i := 1; i <= max; i++ {
		numbers = append(numbers, i)
	}

	var sumOfSquares int

	for _, i := range numbers {
		sumOfSquares += util.Square(i)
	}

	squareOfSum := util.Square(util.Sum(numbers))

	fmt.Println(squareOfSum - sumOfSquares)
}
