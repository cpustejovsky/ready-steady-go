package main

import (
	"fmt"

	"github.com/cpustejovsky/ready-steady-go/tree/master/GoPrgrammingLanguage/ch2/popcount"
)

func main() {
	fmt.Println(popcount.PopCount(0))
	fmt.Println(popcount.PopCountLoop(0))
	fmt.Println(popcount.PopCount(1))
	fmt.Println(popcount.PopCountLoop(1))
	fmt.Println(popcount.PopCount(10))
	fmt.Println(popcount.PopCountLoop(10))
	fmt.Println(popcount.PopCount(6))
	fmt.Println(popcount.PopCountLoop(6))
}
