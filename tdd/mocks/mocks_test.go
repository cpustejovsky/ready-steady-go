package mocks_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/cpustejovsky/ready-steady-go/tdd/mocks"
)

const finalWord = "Go!"
const countdownStart = 3

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	s := &SpySleeper{}
	mocks.Countdown(s, countdownStart, finalWord, buffer)

	got := buffer.String()
	want := `3
2
1
Go!
`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if s.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", s.Calls)
	}

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		mocks.Countdown(spySleepPrinter, countdownStart, finalWord, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}
