package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	val map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) inc(key string) {
	c.mux.Lock()
	// Only one goroutine at a time
	// can access the map[string]int
	c.val[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) getValueBy(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	// Use defer to ensure the mutex will be unlocked
	// after function return result
	return c.val[key]
}

func main() {
	counter := SafeCounter{val: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		go counter.inc("lalala")
	}

	time.Sleep(time.Second)
	fmt.Println(counter.getValueBy("lalala"))
}
