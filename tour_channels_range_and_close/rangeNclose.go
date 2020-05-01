package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		// send x to channel c
		c <- x

		x = y
		y = x + y
	}

	// close channel c
	close(c)
}

func main() {
	c := make(chan int, 10)

	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}

}
