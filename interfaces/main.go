package main

import "fmt"

// if a type has a function called getGreeting() that returns a string,
// then it is a bot and can use functions that take a bot type
type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
	printGreetingLength(eb)
	printGreetingLength(sb)
}

func (eb englishBot) getGreeting() string {
	//VERY custom logic for generating an English greeting
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	//VERY custom logic for generating an Spanish greeting
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func printGreetingLength(b bot) {
	fmt.Println(len(b.getGreeting()))
}
