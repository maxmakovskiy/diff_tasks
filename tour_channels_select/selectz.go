package main

import "fmt"

func fibonacci(ch, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case ch <- x:
			x = y
			y = x + y
		case <-quit:
			fmt.Println("QUIT")
			return
		}
	}
}

func main() {

	ch := make(chan int)
	quitCh := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quitCh <- 0
	}()

	fibonacci(ch, quitCh)

}
