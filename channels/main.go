package main

import (
	"fmt"
	"net/http"
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

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	fmt.Println(<-c)
	fmt.Println("all done!")
}

func checkLink(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Printf("There's an error:\n%v", err)
		c <- "Error Occurred"
	}
	if resp.StatusCode == 200 {
		fmt.Printf("%v is working.\n Status Code %v\n", link, resp.StatusCode)
		c <- "should be working"
	} else {
		fmt.Printf("%v is not working.\n Status Code: %v\n", link, resp.StatusCode)
		c <- "should NOT be working"
	}
}
