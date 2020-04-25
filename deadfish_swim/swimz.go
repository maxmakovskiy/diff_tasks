package main

import (
	"fmt"
	"strings"
)

// "i" - incr value
// "d" - decr value
// "s" - sqr value
// "o" - output value
// "iiisdoso" -> []int{8, 64}
func Parse(data string) []int {
	res := make([]int, 0)
	var value int
	for _, k := range strings.Split(data, "") {
		if t := strings.Compare(k, "i"); t == 0 {
			value++
		} else if t := strings.Compare(k, "d"); t == 0 {
			value--
		} else if t := strings.Compare(k, "s"); t == 0 {
			value = (value * value)
		} else if t := strings.Compare(k, "o"); t == 0 {
			res = append(res, value)
		}
	}

	return res
}

func main() {
	src := "iiisdoso"
	fmt.Printf("init value -> %s\nresult -> %v\n", src, Parse(src))
}
