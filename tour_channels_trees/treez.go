package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	var current, pre *tree.Tree

	if t != nil {
		current = t
		for current != nil {
			if current.Left == nil {
				ch <- current.Value
				current = current.Right
			} else {
				pre = current.Left
				for (pre.Right != nil) && (pre.Right != current) {
					pre = pre.Right
				}

				if pre.Right == nil {
					pre.Right = current
					current = current.Left
				} else {
					pre.Right = nil
					ch <- current.Value
					current = current.Right
				}
			}
		}
	}

	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	var str1, str2 string
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v := range ch1 {
		str1 += string(v)
	}

	for v := range ch2 {
		str2 += string(v)
	}

	if str1 == str2 {
		return true
	}

	return false
}

/*
type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}
*/

func main() {
	fmt.Println(Same(tree.New(2), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(1)))

	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
