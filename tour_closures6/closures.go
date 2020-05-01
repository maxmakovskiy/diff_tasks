package main

import "fmt"

func funfunfunctions() []func() {

	arr := []int{1, 2, 3, 4}
	result := make([]func(), 0)

	for i := range arr {
		result = append(result, func() { fmt.Printf("index - %d, value - %d\n", i, arr[i]) })
	}

	fmt.Println(result)

	return result
}

func main() {
	functions := funfunfunctions()

	for f := range functions {
		functions[f]()
	}
}
