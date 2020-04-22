package main

import "fmt"

func PositiveSum(numbers []int) int {
	var res int

	for _, v := range numbers {
		if v > 0 {
			res += v
		}
	}

	return res
}

func main() {
	nums := []int{-1, 2, 3, 4, -5}
	fmt.Printf("SOURCE: %#v\nRESULT: %d", nums, PositiveSum(nums))
}
