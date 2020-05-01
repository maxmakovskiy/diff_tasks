package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	source := "Hello, Reader!"
	strReader := strings.NewReader(source)
	data := make([]byte, 8)

	fmt.Println("SOURCE -> ", source)

	for {
		numOfBytes, err := strReader.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Printf("Number of populated bytes = %v; error = %v\n", numOfBytes, err)
		fmt.Printf("data[:numOfBytes] = %q\n", data[:numOfBytes])
	}
}
