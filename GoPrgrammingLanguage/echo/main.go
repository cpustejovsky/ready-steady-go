package main

import (
	"fmt"
	"os"
	"strings"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func customEcho() {
	var s string
	fmt.Println(os.Args[0])
	for i, arg := range os.Args[1:] {
		s += fmt.Sprintf("index %d) %v\n", i, arg)
	}
	fmt.Println(s)
}
func main() {
	echo1()
	echo2()
	echo3()
	// customEcho()
}
