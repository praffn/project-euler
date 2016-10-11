package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"./util"
)

func main() {
	inputStr, err := ioutil.ReadFile("013-input.txt")
	util.Check(err)
	input := strings.Split(string(inputStr), "\n")

	var numbers []int
	for _, num := range input {
		if len(num) < 1 {
			break
		}
		n, err := strconv.Atoi(num[:11])
		util.Check(err)
		numbers = append(numbers, n)
	}

	sum := util.Sum(numbers)
	sumStr := strconv.Itoa(sum)
	fmt.Println(sumStr[:10])
}
