package main

import "fmt"

// named function
func printMessage(message string) {
	fmt.Println(message)
}

func getPrintMessage() func(string) {
	// return an anonumous function
	return func(message string) {
		fmt.Println(message)
	}
}

func main() {
	// calls named function
	printMessage("Hello function!")

	// anonymous function declared and called here
	func(message string) {
		fmt.Println(message)
	}("Hello anonymous function from main!")

	// get anonymous function and calls it
	printFunc := getPrintMessage()
	printFunc("Hello anonymous function using caller!")
}
