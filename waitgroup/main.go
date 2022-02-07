package main

import (
	"fmt"
	"sync"
)

func waitGroup() {
	begin := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()
}

func JustChannels() {
	begin := make(chan int)
	go func() {
		defer close(begin)
		for i := 0; i < 5; i++ {
			fmt.Printf("%v has begun\n", i)
			begin <- i
		}
	}()
	for v := range begin {
		fmt.Printf("%v has ended\n", v)
	}
}
