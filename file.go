package main

import (
	"io"
	"os"
)

func readSize(f *os.File, size int) ([]byte, error) {
	buffer := make([]byte, size)
	n, err := f.Read(buffer)
	if err == io.EOF {
		return nil, nil
	}
	return buffer[0:n], err
}
