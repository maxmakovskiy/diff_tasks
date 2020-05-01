package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(source byte) byte {
	out := source

	if source > 64 && source < 91 {
		out = source - 64
		out = out + 13
		out = out % 26
		out = out + 64
	}
	if source > 96 && source < 123 {
		out = source - 96
		out = out + 13
		out = out % 26
		out = out + 96
	}

	return out
}

func (reader rot13Reader) Read(b []byte) (int, error) {
	for i, v := range b {
		b[i] = rot13(v)
	}

	num, err := reader.r.Read(b)
	return num, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
