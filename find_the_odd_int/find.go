package main

import "fmt"

// find the integer that appears
//   an odd number of times
func findOdd(seq []int) int {
	var tempMap map[int]int = make(map[int]int)
	var result int

	for i, _ := range seq {
		tempMap[seq[i]] += 1
	}

	fmt.Println(tempMap)

	for k, v := range tempMap {
		if v%2 != 0 {
			result = k
		}
	}

	return result
}

func main() {
	arr := []int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}
	fmt.Println(findOdd(arr))
}
