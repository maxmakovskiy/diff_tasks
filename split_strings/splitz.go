package main

import (
	"fmt"
	"regexp"
)

func divideString(str string) []string {
	re := regexp.MustCompile("..")
	res := re.FindAllString(str, -1)
	return res
}

func Solution(str string) (res []string) {
	if len(str)%2 == 0 {
		res = divideString(str)
	} else {
		res = divideString(str + "_")
	}

	return
}

func main() {
	fmt.Println(divideString("awsaws"))
}
