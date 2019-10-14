package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	data := httpFetcher("https://www.cpustejovsky.com")
	fmt.Printf("%s", data)
	saveToFile("test.html", data)
}

func httpFetcher(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Got a nil\n here's the error:\n%v", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func saveToFile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0666)
}