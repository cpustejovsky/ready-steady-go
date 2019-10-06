package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"Hi", "there", "how", "are", "you"}
	fmt.Println(mySlice)
	updateSlice(mySlice)
	fmt.Println(mySlice)

	myString := "Test"
	fmt.Println(myString)
	updateString(&myString)
	fmt.Println(myString)

}

func updateSlice(s []string) {
	s[0] = "Howdy"
}

func updateString(s *string) {
	*s = "Testing complete"
}
