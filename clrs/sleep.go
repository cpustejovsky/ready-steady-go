package clrs

import (
	"fmt"
	"sync"
	"time"
)

func SleepSort(arr []int) {
	var wg sync.WaitGroup
	for _, num := range arr {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Println(n)
		}(num)
	}
	wg.Wait()
}
