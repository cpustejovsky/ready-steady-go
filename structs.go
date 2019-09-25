package main

import "fmt"

type Vertex struct {
	X, Y int
}

type person struct { 
	name string
	age int 
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	v4 = Vertex{Y: 2}
	p  = &Vertex{1, 2} // has type *Vertex
	cpustejovsky  = person{name: "Charles", age: 26}
)

func main() {
	fmt.Println(v1, p, v2, v3, v4, cpustejovsky.name)
}
