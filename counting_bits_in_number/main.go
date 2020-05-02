package main

import "fmt"

// "Elements of Programming Interviews (C++)"
// paragraph "Primitive types"
func countBits(num int) (res int) {
	// num -> 5 -> 0101
	for num != 0 {
		// 0101 & 0001 = 0001
		res += num & 1
		// 0101 >> 1 = 0010
		num >>= 1
		// and so on
	}

	return
}

func main() {
	var num int = 5 // 0101 -> 2 bits
	fmt.Printf("Number of bits in %d equals to %d\n", num, countBits(num))
}
