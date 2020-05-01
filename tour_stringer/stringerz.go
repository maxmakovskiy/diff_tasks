package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func (v Vertex) String() string {
	return fmt.Sprintf("(%d;%d)", v.X, v.Y)
}

func main() {
	v1 := Vertex{2, 5}
	v2 := Vertex{-2, -1}

	fmt.Println(v1, v2)

}
