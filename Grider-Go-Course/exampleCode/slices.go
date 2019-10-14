package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range nums {
		if num == 0 {
			fmt.Printf("%v is zero\n", num)
		}
		if num%2 == 0 && num != 0 {
			fmt.Printf("%v is even\n", num)
		}
		if num%2 == 1 {
			fmt.Printf("%v is odd\n", num)
		}
	}
}
