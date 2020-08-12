package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	// resp, err := http.Get("https://www.cpustejovsky.com")
	// if err != nil {
	// 	fmt.Printf("Got a nil\n here's the error:\n%v", err)
	// 	os.Exit(1)
	// }

	// lw := logWriter{}

	// io.Copy(lw, resp.Body)

	data := httpFetcher("https://www.cpustejovsky.com")
	saveToFile("test.html", data)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Just wrote this many bytes: %v\n", len(bs))
	return len(bs), nil
}

func saveToFile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0666)
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
