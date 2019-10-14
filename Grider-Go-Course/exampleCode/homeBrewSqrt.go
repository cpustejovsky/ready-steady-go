package main

import (
	"fmt"
	"math"
)

func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func Sqrt(x float64) float64 {
	z := 1.0
	count := 0
	for count < 10 {
		z -= (z*z - x) / (2*z)
		y := z*z
		if Round(y, 0.00001) == x {
			return z;
		}
		count++
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}