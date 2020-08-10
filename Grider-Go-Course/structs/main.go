package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	age       int
	married   bool
	contactInfo
}

/*
	Turn an address into a value with *address
	Turn a value into an address with %value
*/

func main() {
	cpustejovsky := person{firstName: "Chas", lastName: "Pustejovsky", age: 27, married: true}
	cpustejovsky.contactInfo.email = "charles.pustejovsky@gmail.com"
	cpustejovsky.contactInfo.zip = 30004

	catherine := person{firstName: "Catherine", lastName: "Pustejovsky", age: 25, married: true}
	/*
	  "&variable" operator:
	  gives the memory address of the value the variable is pointing at
	*/
	cpustejovskyPointer := &cpustejovsky
	cpustejovsky.print()
	cpustejovskyPointer.updateName("Chazz")
	cpustejovsky.print()
	cpustejovsky.simpleUpdateName("Charles")
	cpustejovsky.print()
	catherine.print()
}

/*
	*person provides a type description:
	This means we're working with a pointer to a person type
*/
func (pointerToPerson *person) updateName(newFirstName string) {
	/*
		This is an operator:
		it means we want to manipulate the value the pointer is referencing
	*/
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson person) print() {
	fmt.Printf("%+v\n", pointerToPerson)
}

/*
	You don't need to massage pointer out of person with "&variable"
	Here is the syntatic sugar
*/
func (pointerToPerson *person) simpleUpdateName(newFirstName string) {
	(pointerToPerson).firstName = newFirstName
}
