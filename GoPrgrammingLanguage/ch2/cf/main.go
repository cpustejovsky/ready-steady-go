package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/cpustejovsky/ready-steady-go/tree/master/GoPrgrammingLanguage/ch2/tempconv"
)

func main() {
	var kind string
	flag.StringVar(&kind, "kind", "default", "the kind of conversion you wish to calculate")
	flag.Parse()

	for i, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil && i > 1 {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		switch kind {
		case "length":
			fmt.Println("length goes here")
		case "temperature":
			temperatureConv(t)
		case "default":
			fmt.Println("default is running as temperature")
			temperatureConv(t)
		}
	}
}

func temperatureConv(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}
