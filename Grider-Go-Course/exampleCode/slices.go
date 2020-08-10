package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range nums {
		switch {
		case num == 0:
			fmt.Printf("%v is zero\n", num)
		case num%2 == 0 && num != 0:
			fmt.Printf("%v is even\n", num)
		case num%2 == 1:
			fmt.Printf("%v is odd\n", num)
		}
	}
}
