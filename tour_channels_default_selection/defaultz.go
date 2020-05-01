package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	boom := time.After(5 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("Tick!")
		case <-boom:
			fmt.Println("BOOM!!!")
			return
		default:
			fmt.Println("wait...")
			time.Sleep(500 * time.Millisecond)
		}
	}

}
