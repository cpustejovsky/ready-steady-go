package main

import (
	"fmt"
)

func spacer() {
	fmt.Println("=====================================================================")
}

func main() {
	deck := newDeckFromFile("test")
	// deck.shuffle()
	fmt.Println(len(deck))
	fmt.Println(deck[0])
	fmt.Println(deck[len(deck)-1])
}
