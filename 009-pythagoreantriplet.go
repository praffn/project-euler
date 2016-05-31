// A Pythagorean triplet is a set of three natural numbers,a < b < c, for which,
//
//                              a^2 + b^2 = c^2
//
// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
//
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.
package main

import (
	"fmt"

	"./util"
)

func findProduct(n int) (p int) {
	var a, b, c int

	for a = 1; a < n; a++ {
		for b = 2; b <= n; b++ {
			if a < b {
				c = n - b - a
				if c > b && util.Square(a)+util.Square(b) == util.Square(c) {
					p = a * c * b
					break
				}
			}
		}
	}

	return p
}

func main() {
	fmt.Println(findProduct(1000))
}
