//-----------------------------------------------------------------------------
package main

import "fmt"

// Create a new type of "deck" which is a slice of strings

type deck []string

//receiver function: any variable of type ___ gets access to the method defined
//by convention, the variable has the first letter of the type, so "d" in case 
//of type "deck"
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
