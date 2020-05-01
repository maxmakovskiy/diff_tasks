package main

import "fmt"

func outer(message string) func() {
	text := "Modified " + message

	// closure
	// function has access to text even
	//   after exiting this block
	foo := func() {
		fmt.Println(text)
	}

	// return the closure
	return foo
}

func main() {
	// foo is a closure
	foo := outer("args")

	// calling a closure
	foo()
}
