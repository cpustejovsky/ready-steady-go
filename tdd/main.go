package main

import (
	"os"
	"time"

	"github.com/cpustejovsky/ready-steady-go/tdd/mocks"
)

const finalWord = "Blastoff!"
const countdownStart = 4

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	dt := &DefaultSleeper{}
	mocks.Countdown(dt, countdownStart, finalWord, os.Stdout)
}
