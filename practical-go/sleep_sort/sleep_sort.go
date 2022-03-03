package main

import (
	"fmt"
	"time"
)

func main() {
	nums := []int{23, 8, 15, 4, 42, 16}
	fmt.Println(sleepSort(nums)) // 4 8 15 16 23 42
}

/* Algorithm:
For every number "n" in values, spin a goroutine that will:
- Sleep n milliseconds
- Send n back over a channel

In the main function, collect values back from the channel

NEVER USE IRL!
*/

func sleepSort(values []int) []int {
	c := make(chan int, len(values))
	for _, v := range values {
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Millisecond)
			c <- n
		}(v)
	}
	var out []int
	for range values {
		out = append(out, <-c)
	}
	return out
}

//I-O bound vs CPU bound
//spinning more goroutines helps for IO (see scheduling blog)
//more processors for CPU bound? TODO: find out!
