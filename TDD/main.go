package main

import (
	"os"
	"time"

	"github.com/cpustejovsky/ready-steady-go/mocking"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	cs := &mocking.ConfigurableSleeper{1 * time.Second, time.Sleep}
	mocking.Countdown(os.Stdout, cs)
}
