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
	now := time.Now()
	doneTime := now.Add(time.Minute * time.Duration(total))
	endOfday := time.Date(now.Year(), now.Month(), now.Day(), 21, 0, 0, 0, now.Location())
	freeTime := endOfday.Sub(doneTime)
	fmt.Printf("TODO Time: %.2fhr (%.2fmin) remaining \nYou have %.2fhr (%.2fmin) of free time\n", hours, total,
		freeTime.Hours(), freeTime.Minutes())
}
