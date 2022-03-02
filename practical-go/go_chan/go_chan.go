package main

/*
Channel semantics:
- Send/Receive to/from a channel will block *
- Reveise from a closed channel won't block and return the zero value for the channel type
- Closing a chann will free all receiving goroutine
- Send to a closed channel will PANIK!
- Closing a closed channel will PANIK!
- Receive/send from a nil channel will block forever

*/
import (
	"fmt"
	"time"
)

func main() {
	/*
		General notes!
		When working with gooroutines, ask "What is being shared ?""
		You have no control of a goroutine once you spin one up so collaboration is key
		Not thread safe
		Locks and semaphores are difficult to get right
		Global locking is slow; fine-grained locking is prone to bugs
		"Don't communicate by sharing memory, share memory by communicating." - Rob Pike
	*/
	// for i := 0; i < 3; i++ {
	// bug because of function closures
	// go func() {
	// 	fmt.Println(i)
	// }()
	// FIX 1: pass in param
	// go func(n int) {
	// 	fmt.Println(n)
	// }(i)
	// FIX 2: New local variable
	// i := i // in this line shadows line 12
	// go func() {
	// 	fmt.Println(i)
	// }()
	// }

	num := 3
	ch := make(chan int, num)
	go func() {
		defer close(ch)
		for i := 0; i < num; i++ {
			ch <- i
		}
	}()
	// syntatic sugar for a while loop that uses val, ok with breaks
	for v := range ch {
		fmt.Println(v)
	}
	//Reading from closed channel gets you the zero value
	val, ok := <-ch
	fmt.Println(val)
	//ok lets you know if it's closed
	//val, ok paradigm is used for type assertions and maps as well
	fmt.Println(ok)

	// ch <- 3 // WILL PANIC

	//Channels can be used for synchronization
	/*
		Known as a barrier in another language
		Used to help with waiting on something.
		When the data comes in you, close the signal/sync channel

	*/
	sync := make(chan struct{})
	for i := 0; i < num; i++ {
		i := i
		go func() {
			fmt.Printf("%d is waiting\n", i)
			<-sync
			fmt.Printf("%d is freed!\n", i)

		}()
	}
	time.Sleep(time.Millisecond)
	fmt.Println("Who let the goroutines out?")
	close(sync)
	time.Sleep(time.Millisecond)

	var nilch chan int

}
