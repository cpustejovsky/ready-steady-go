package mocks

import (
	"fmt"
	"io"
)

type Sleeper interface {
	Sleep()
}

func Countdown(s Sleeper, start int, finalWord string, w io.Writer) {
	for i := start; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}

	s.Sleep()
	fmt.Fprintln(w, finalWord)
}
