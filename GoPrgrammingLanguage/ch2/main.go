package main

import (
	"fmt"

	"github.com/cpustejovsky/ready-steady-go/tree/master/GoPrgrammingLanguage/ch2/conv"
)

func main() {
	fmt.Printf("Brrrr!	%v\n", conv.AbsoluteZeroC)
	fmt.Println(conv.CToF(conv.BoilingC))
	fmt.Println(conv.CToK(conv.BoilingC))
}
