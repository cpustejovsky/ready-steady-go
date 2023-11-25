package clrs

import (
	"sync"
	"time"
)

func SleepSort(arr []int) []int {
	var wg sync.WaitGroup
	var nums = make(chan int, len(arr))
	var sorted []int
	for _, num := range arr {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Millisecond)
			nums <- n
		}(num)
	}
	wg.Wait()
	close(nums)
	for num := range nums {
		sorted = append(sorted, num)
	}
	return sorted
}
