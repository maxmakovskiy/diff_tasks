package main

import "fmt"

func sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	// send sum to ch
	ch <- sum
}

func main() {
	s := []int{2, 3, 4, 1, 2, -22, 3, 4, 1, 2, -2}

	c := make(chan int)

	go sum(s, c)
	go sum(s[len(s)/2:], c)

	// receive from c
	// result1, result2 := <-c, <-c
	result1 := <-c
	result2 := <-c

	fmt.Printf("Sum of %v equals %d\n", s, result1)
	fmt.Printf("Sum of %v eqauls %d\n", s[len(s)/2:], result2)
}
