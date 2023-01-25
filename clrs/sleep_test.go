package clrs_test

import (
	"github.com/cpustejovsky/ready-steady-go/clrs"
	"testing"
)

func TestSleepSort(t *testing.T) {
	input := []int{5, 2, 4, 6, 1, 3, 10}
	clrs.SleepSort(input)
}
