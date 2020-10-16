package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	key := make([]byte, 32)
	fmt.Println(key)
	_, err := rand.Read(key)
	if err != nil {
		// handle error here
	}
	fmt.Println(string(key[:]))
}
