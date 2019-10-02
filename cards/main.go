package main

import "fmt"

func spacer() {
	fmt.Println("=====================================================================")
}

func main() {
	// cards := newDeck()
	// cards.print()
	// spacer()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// spacer()
	// remainingDeck.print()
	// stringCards := cards.deckToString()
	// fmt.Println(stringCards)
	// cards.saveToFile("test")
	deck := newDeckFromFile("test")
	fmt.Println(deck)
}
