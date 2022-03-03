/*
Write a function that gets an index file with names of files and sha256
signatures in the following format
0c4ccc63a912bbd6d45174251415c089522e5c0e75286794ab1f86cb8e2561fd  taxi-01.csv
f427b5880e9164ec1e6cda53aa4b2d1f1e470da973e5b51748c806ea5c57cbdf  taxi-02.csv
4e251e9e98c5cb7be8b34adfcb46cc806a4ef5ec8c95ba9aac5ff81449fc630c  taxi-03.csv
...

You should compute concurrently sha256 signatures of these files and see if
they math the ones in the index file.

- Print the number of processed files
- If there's a mismatch, print the offending file(s) and exit the program with
  non-zero value

Get taxi.zip from the web site and open it. The index file is sha256sum.txt

Bonus:
- Used a fixed pool of n goroutines (say 5)
- Make sure to tell the worker there's no more work
- Tell workers to terminate once there's an error
*/
package main

import (
	"bufio"
	"compress/bzip2"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

func fileSig(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, bzip2.NewReader(file))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// Parse signature file. Return map of path->signature
func parseSigFile(r io.Reader) (map[string]string, error) {
	sigs := make(map[string]string)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		// Line example
		// 6c6427da7893932731901035edbb9214  nasa-00.log
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			// TODO: line number
			return nil, fmt.Errorf("bad line: %q", scanner.Text())
		}
		sigs[fields[1]] = fields[0]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sigs, nil
}

type result struct {
	fileName string
	err      error
	match    bool
}

//sigWorker is just a normal function return the result value;
func sigWorker(fileName string, expected string) result {
	r := result{
		fileName: fileName,
	}
	sig, err := fileSig(fileName)
	r.err = err
	r.match = sig == expected
	return r
}

func main() {
	rootDir := "./taxi"
	file, err := os.Open(path.Join(rootDir, "sha256sum.txt"))
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	sigs, err := parseSigFile(file)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	start := time.Now()

	ch := make(chan result)
	//this is how you determine your processes in Go
	g := runtime.GOMAXPROCS(0)
	//creates a buffered channel that will block when the buffer is full
	sem := make(chan struct{}, g)
	for name, signature := range sigs {
		fileName := path.Join(rootDir, name) + ".bz2"
		//sem is helping us create a semaphore pattern to only allow g workers at a time where g is the number of cores on our machine
		go func(f string, s string) {
			//sem receives value
			sem <- struct{}{}
			// work is done and sent to result channel
			ch <- sigWorker(f, s)
			// sem sends value, freeing up space so another value can be received and more work can continue
			<-sem
		}(fileName, signature)
	}

	for range sigs {
		r := <-ch
		if r.err != nil {
			fmt.Fprintf(os.Stderr, "error: %s - %s\n", r.fileName, r.err)
			continue
		}

		if !r.match {
			fmt.Printf("%s: mismatch\n", r.fileName)
		}
	}

	duration := time.Since(start)
	fmt.Printf("concurrently processed %d files in %v\n", len(sigs), duration)

	start2 := time.Now()
	for name, signature := range sigs {
		fileName := path.Join(rootDir, name) + ".bz2"
		result := sigWorker(fileName, signature)
		if result.err != nil {
			fmt.Fprintf(os.Stderr, "error: %s - %s\n", result.fileName, result.err)
			os.Exit(1)
		}
		if !result.match {
			fmt.Printf("%s: mismatch\n", result.fileName)
		}
	}

	duration2 := time.Since(start2)
	fmt.Printf("sequentially processed %d files in %v\n", len(sigs), duration2)

}
