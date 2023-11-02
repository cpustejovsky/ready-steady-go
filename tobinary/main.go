package main

import (
	"fmt"
	"os"
	"strconv"
)

func Convert(nums ...int) string {
	var out string
	for _, n := range nums {
		out += fmt.Sprintf("%d is %b in binary_search\n", n, n)
	}
	return out
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please provide number")
		os.Exit(1)
	}
	var nums []int
	for _, num := range os.Args[1:] {
		n, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		nums = append(nums, n)
	}
	fmt.Println(Convert(nums...))
}
