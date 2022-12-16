package main

import (
	"fmt"
	"math"
	"strconv"
)

func inflation(starting, rate, years float64) float64 {
	return starting * math.Pow(1+rate, years)
}

func fizzBuzz(n int) []string {
	var arr []string
	for i := 1; i <= n; i++ {
		switch {
		case i % 3 == 0 && i % 5 == 0:
			arr = append(arr, "FizzBuzz")
		case i % 3 == 0:
			arr = append(arr, "Fizz")
		case i % 5 == 0:
			arr = append(arr, "Buzz")
		default:
			arr = append(arr, strconv.Itoa(i))
		}
	}
	return arr
}

func twoSum(nums []int, sum int) []int {
	var result []int
	for i, augend := range nums {
		for j, addend := range nums {
			if augend + addend == sum && i != j {
				result = []int{i, j}
				return result
			}
		}
	}
	return result
}

func main() {
	fmt.Println(twoSum([]int{2,4,11,3}, 6))
}
