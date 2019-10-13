package main

import (
	"fmt"
	"os"
)

func main() {
	argv := os.Args[1]
	file, err := os.Open(argv)
	if err != nil {
		fmt.Printf("Got a nil\n here's the error:\n%v", err)
		os.Exit(1)
	}
	data := make([]byte, 32*1024)
	count, err := file.Read(data)
	fmt.Print(string(data[:count]))

}
