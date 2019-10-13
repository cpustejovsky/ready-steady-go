package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	argv := os.Args[1]
	file, err := os.Open(argv)
	if err != nil {
		fmt.Printf("Got a nil\n here's the error:\n%v", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, file)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Just wrote this many bytes: %v\n", len(bs))
	return len(bs), nil
}
