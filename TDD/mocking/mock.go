package mocking

import (
	"fmt"
	"io"
	"time"
)

type ConfigurableSleeper struct {
	Duration    time.Duration
	SleeperFunc func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.SleeperFunc(c.Duration)
}

type Sleeper interface {
	Sleep()
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(out, i)
	}
	s.Sleep()
	fmt.Fprint(out, finalWord)
}
