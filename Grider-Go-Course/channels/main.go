package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
		"https://bitpay.com/resources",
	}
	// Too many channels will make your program hang up
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Printf("There's an error:\n%v", err)
		c <- link
	}
	if resp.StatusCode == 200 {
		fmt.Printf("%v is working.\nStatus Code %v\n", link, resp.StatusCode)
		c <- link
	} else {
		fmt.Printf("%v is not working.\nStatus Code: %v\n", link, resp.StatusCode)
		c <- link
	}
}
