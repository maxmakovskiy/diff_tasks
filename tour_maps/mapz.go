package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	var result map[string]int = make(map[string]int)

	for _, word := range strings.Fields(s) {
		//fmt.Println(word)
		result[word] += 1
	}

	return result
}

func main() {
	var s string = "Lorem ipsum dolor sit amet amet sit dolor dolor dolor"
	fmt.Println("Source -> ", s, "\nResult -> ", WordCount(s))

}
