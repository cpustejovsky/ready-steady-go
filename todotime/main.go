package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("please provide minutes")
		os.Exit(1)
	}
	var total float64
	for _, duration := range os.Args[1:] {
		m, err := strconv.ParseFloat(duration, 0)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		total += m
	}

	hours := total / 60
	doneTime := time.Now().Add(time.Minute * time.Duration(total))
	fmt.Printf("TODO Time: %.2fhr (%.2fmin) remaining \nSoonest time to finish: %s\n", hours, total,
		doneTime.Format(time.TimeOnly))
}
