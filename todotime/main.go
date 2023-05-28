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
	minutes, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	hours := minutes / 60
	remainder := minutes - hours*60
	doneTime := time.Now().Add(time.Minute * time.Duration(minutes))
	fmt.Printf("the soonest you can finish your To-Dos (%dhr %dmin) is %v\n", hours, remainder, doneTime.Format(time.TimeOnly))
}
