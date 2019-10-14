package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	base   float64
	height float64
}

func main() {
	tri := triangle{
		base:   3.0,
		height: 7.0,
	}
	squ := square{
		sideLength: 7.0,
	}
	printArea(squ)
	printArea(tri)

}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Printf("Here is the area! %v\n", s.getArea())
}
