package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

func inflation(starting, rate, years float64) float64 {
	return starting * math.Pow(1+rate, years)
}

func fizzBuzz(n int) []string {
	var arr []string
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			arr = append(arr, "FizzBuzz")
		case i%3 == 0:
			arr = append(arr, "Fizz")
		case i%5 == 0:
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
			if augend+addend == sum && i != j {
				result = []int{i, j}
				return result
			}
		}
	}
	return result
}

func removeDigits(r rune) rune {
	if unicode.IsDigit(r) {
		r = -1
	}
	return r
}

//Input := "H1e5ll0o"
//fmt.Println(Input)
//Output := strings.Map(removeDigits, Input)
//fmt.Println(Output) // "Hello"

func Reverse(nums []int) []int {
	j := len(nums) - 1
	for i := 0; i < j; i++ {
		fmt.Printf("index i: %d; index j: %d\n", i, j)
		nums[i], nums[j] = nums[j], nums[i]
		j--
	}

	return nums
}

func main() {
	initial := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(Reverse(initial))
}
