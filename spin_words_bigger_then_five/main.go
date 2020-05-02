package main

import (
	"fmt"
	"strings"
)

func reverse(source string) string {
	result := []rune(source)

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		temp := result[i]
		result[i] = result[j]
		result[j] = temp
	}

	return string(result)
}

// reverse word when len(word) >= 5
func SpinWords(str string) string {
	var result string

	temp := strings.Split(str, " ")
	for i, v := range temp {
		if (v == " ") || (len(v) < 5) {
			if i != len(temp) {
				result += (v + " ")
			} else {
				result += v
			}
		} else {
			if i != len(temp) {
				result += (reverse(v) + " ")
			} else {
				result += reverse(v)
			}
		}
	}

	return result
}

func main() {
	str := "Hi pizza"
	fmt.Println(SpinWords(str)) // expected "Hi azzip"
}
