package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct {
}

func (reader MyReader) Read(data []byte) (int, error) {
	data[0] = byte(int('A'))
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
