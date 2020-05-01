package main

import (
	"fmt"
	"time"
)

func funcOne(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "one"
}

func funcTwo(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "two"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go funcOne(ch1)
	go funcTwo(ch2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch2:
			fmt.Println("received", msg1)
		case msg2 := <-ch1:
			fmt.Println("received", msg2)
		}
	}
}
