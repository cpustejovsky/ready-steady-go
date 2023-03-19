package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		log.Println("please provide denominator and then numerator")
		os.Exit(1)
	}
	dint, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	nint, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	numerator := float64(nint)
	denominator := float64(dint)
	percentage := numerator / denominator * 100
	roundedPerc := math.Round(percentage*100) / 100
	fmt.Printf("\n%v%s Complete!\n", roundedPerc, "%")
}
