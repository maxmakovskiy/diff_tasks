package main

import "fmt"

func outer(message string) {
	// variable in outer function
	text := "Modified " + message

	// foo is a closure
	// foo us a inner function and has access to <text> variable
	// closures have access to the variables
	//   even after exiting this block
	foo := func() {
		fmt.Println(text)
	}

	// calling the closure
	foo()
}

func main() {
	outer("arguments")
}
