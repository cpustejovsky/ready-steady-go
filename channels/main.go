package main

import (
	"fmt"
	"runtime"
)

func main() {
	g := runtime.GOMAXPROCS(4)
	fmt.Println(g)
}
