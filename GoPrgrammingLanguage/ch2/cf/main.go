package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/cpustejovsky/ready-steady-go/tree/master/GoPrgrammingLanguage/ch2/conv"
)

func main() {
	var kind string
	flag.StringVar(&kind, "kind", "default", "the kind of conversion you wish to calculate")
	flag.Parse()

	for i, arg := range os.Args[1:] {
		m, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			if i == 0 {
				continue
			} else {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
		}
		switch kind {
		case "length":
			lengthConv(m)
		case "temperature":
			temperatureConv(m)
		case "default":
			fmt.Println("default is running as temperature")
			temperatureConv(m)
		}
	}
}

func temperatureConv(t float64) {
	f := conv.Fahrenheit(t)
	c := conv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", f, conv.FToC(f), c, conv.CToF(c))
}

func lengthConv(l float64) {
	ft := conv.Foot(l)
	m := conv.Meter(l)
	fmt.Printf("%s = %s, %s = %s\n", ft, conv.FtToM(ft), m, conv.MToFt(m))
}
