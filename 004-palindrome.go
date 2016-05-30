// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//
// Find the largest palindrome made from the product of two 3-digit numbers.
package main

import (
	"fmt"
	"strconv"

	"./util"
)

func main() {
	var p int
	max := 100001
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			p = i * j
			if max < p && util.IsPalindrome(strconv.Itoa(p)) {
				max = p
			}
		}
	}
	fmt.Println(max)
}
