package main

import "fmt"

func counter(start int) (func() int, func()) {
	// if the value gets mutated,
	//   the same is reflected in closure
	ctr := func() int {
		return start
	}

	incr := func() {
		start++
	}

	// both ctr and incr have same reference to <start>
	return ctr, incr
}

func main() {
	// creating the one state
	ctr, incr := counter(100)
	// creating the another state
	ctr1, incr1 := counter(100)

	fmt.Println("counter - ", ctr())
	fmt.Println("counter1 - ", ctr1())

	incr()
	fmt.Println("incr()")

	fmt.Println("counter - ", ctr())
	fmt.Println("counter1 - ", ctr1())

	incr1()
	incr1()
	fmt.Println("incr1()x2")

	fmt.Println("counter - ", ctr())
	fmt.Println("counter1 - ", ctr1())

}
