package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = byte(65)
	}
	n = len(b)
	err = nil
	return
}

func main() {
	reader.Validate(MyReader{})
}
