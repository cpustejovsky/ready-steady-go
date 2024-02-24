package main

import "fmt"

func main() {

	funcOne := func(x int) int {
		return x - 2
	}
	funcTwo := func(x int) int {
		return (x*x + x - 6) / (x + 3)
	}
	for i := -100; i <= 100; i++ {
		fmt.Printf("answer for x -2 when x = %d\n%d\n", i, funcOne(i))
		fmt.Printf("answer for (x*x + x - 6)/(x + 3) when x = %d\n%d\n", i, funcTwo(i))
	}
}
