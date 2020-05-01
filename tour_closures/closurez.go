package main

import (
	"fmt"
)

func fibonacci() func() int {
	var sum int
	var a, b int = 0, 1

	return func() int {
		sum = a + b
		a = b
		b = sum

		return sum
	}
}

func main() {
	a := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(a())
	}
}
