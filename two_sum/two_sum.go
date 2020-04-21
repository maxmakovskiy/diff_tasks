package main

import "fmt"

// numbers[a] + numbers[b] = target
// numbers[a] != numbers[b]
// func returns (indexA, indexB)
func TwoSum(numbers []int, target int) [2]int {
	l := len(numbers)

	if l == 2 {
		// if (numbers[0] != numbers[1]) && (numbers[0]+numbers[1] == target) {
		if numbers[0]+numbers[1] == target {
			return [2]int{0, 1}
		} else {
			return [2]int{}
		}
	}

	for i := (l - 1); i >= 0; i-- {
		for j, v := range numbers {
			if v+numbers[i] == target {
				return [2]int{j, i}
			}
		}
	}

	return [2]int{}
}

func main() {
	src1 := []int{1, 2, 3}
	target1 := 4
	fmt.Printf("numbers = %v, target = %d\n --> result of TwoSum(numbers, target)=%v\n", src1, target1, TwoSum(src1, target1))

	src2 := []int{1234, 5678, 9012}
	target2 := 14690
	fmt.Printf("numbers = %v, target = %d\n --> result of TwoSum(numbers, target)=%v\n", src2, target2, TwoSum(src2, target2))

	src3 := []int{2, 2, 3}
	target3 := 4
	fmt.Printf("numbers = %v, target = %d\n --> result of TwoSum(numbers, target)=%v\n", src3, target3, TwoSum(src3, target3))

}
