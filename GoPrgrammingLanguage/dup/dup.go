package main

import (
	"bufio"
	"fmt"
	"os"
)

type dups struct {
	count     int
	locations []string
}

type counts map[string]dups

func main() {
	counts := make(counts)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
		for line, dups := range counts {
			if dups.count > 1 {
				fmt.Printf("%s found %d times in the following files:\t %v\n", line, dups.count, dups.locations)
			}
		}
	}
}

//TODO: come up with a better way than this copying system
func countLines(f *os.File, counts counts) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		c := counts[input.Text()]
		if !includes(c.locations, f.Name()) {
			c.locations = append(c.locations, f.Name())
		}
		c.count++
		counts[input.Text()] = c
	}
}

func includes(array []string, str string) bool {
	for _, value := range array {
		if value == str {
			return true
		}
	}
	return false
}
