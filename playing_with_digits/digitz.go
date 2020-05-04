package main

import (
	"fmt"
	"strconv"
)

func pow(base, exponent int) int {
	var result int = 1

	for i := 1; i <= exponent; i++ {
		result *= base
	}

	return result
}

func getDigits(num int) []int {
	result := make([]int, 0)
	tempString := strconv.Itoa(num)

	for _, c := range tempString {
		val, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Print(err)
		}

		result = append(result, val)
	}

	return result
}

// (p, n)
// (894, 2) | p -> 894, n -> 2
// temp = 8^2 + 9^(2+1) + 4^(2+2) = 8^n + 9^(n+1) + 9^(n+2)
// k = temp/n
// if k is int -> return k
// else return -1
func DigitPow(n, p int) int {
	digits := getDigits(n)

	var temp int
	var degree int = p
	for i, v := range digits {
		temp += pow(v, (p + i))
		degree++
	}

	var k interface{} = temp / n
	if valOfK, _ := k.(int); valOfK != 0 {
		return valOfK
	}

	return -1
}

func main() {
	n := 92
	p := 1
	fmt.Printf("Input = %d; digPow = %d\n", n, DigitPow(n, p))
}
