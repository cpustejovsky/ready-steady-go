package main

import (
	"fmt"
	"time"
)

type heart struct {
	done      <-chan any
	heartbeat chan any
	pulse     <-chan time.Time
	results   chan time.Time
	workGen   <-chan time.Time
}

func New(done <-chan any, pulseInterval time.Duration) *heart {
	pulse := time.Tick(pulseInterval)
	workGen := time.Tick(2 * pulseInterval)
	heartbeat := make(chan any)
	results := make(chan time.Time)
	return &heart{
		done:      done,
		heartbeat: heartbeat,
		pulse:     pulse,
		results:   results,
		workGen:   workGen,
	}
}

func (h *heart) Beat() {
	defer close(h.heartbeat)
	defer close(h.results)
	for {
		select {
		case <-h.done:
			fmt.Println("DONE!")
			return
		case <-h.pulse:
			h.sendPulse()
		case r := <-h.workGen:
			h.sendResult(r)
		}
	}
}

func (h *heart) sendPulse() {
	select {
	case h.heartbeat <- struct{}{}:
	default:
	}
}

func (h *heart) sendResult(r time.Time) {
	for {
		select {
		case <-h.done:
			return
		case <-h.pulse:
			h.sendPulse()
		case h.results <- r:
			return
		}
	}
}

func main() {
	done := make(chan any)
	time.AfterFunc(10*time.Second, func() { close(done) })
	const interval = 1 * time.Second
	h := New(done, interval)
	go h.Beat()
	for {
		select {
		case _, ok := <-h.heartbeat:
			if ok == false {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-h.results:
			if ok == false {
				return
			}
			fmt.Printf("pulse %v\n", r.Format(time.RFC822))
		}
	}
}
