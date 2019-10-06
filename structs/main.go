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
	contact   contactInfo
}

type business struct {
	name     string
	industry string
	contact  contactInfo
}

func main() {
	cpustejovsky := person{firstName: "Charles", lastName: "Pustejovsky", age: 26, married: true}
	cpustejovsky.contact.email = "charles.pustejovsky@gmail.com"
	cpustejovsky.contact.zip = 30004
	fmt.Printf("%+v\n", cpustejovsky)
	var catherine person
	catherine.firstName = "Catherine"
	catherine.lastName = "Pustejovsky"
	catherine.age = 24
	catherine.married = true
	fmt.Printf("%+v\n", catherine)
	bitpay := business{name: "BitPay", industry: "Payment Processing", contact: contactInfo{email: "info@bitpay.com", zip: 30009}}
	fmt.Println(bitpay)
}
