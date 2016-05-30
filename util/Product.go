package util

func Product(nums []int) int {
	product := 1
	for _, v := range nums {
		product *= v
	}
	return product
}
