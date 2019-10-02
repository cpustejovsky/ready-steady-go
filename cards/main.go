package main

import (
	"fmt"
)

func spacer() {
	fmt.Println("=====================================================================")
}

func main() {
	deck := newDeckFromFile("test")
	deck.shuffle()
	fmt.Println(deck)
}
