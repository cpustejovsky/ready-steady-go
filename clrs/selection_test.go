package clrs_test

import (
	"github.com/cpustejovsky/ready-steady-go/clrs"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	input := []int{5, 2, 4, 6, 1, 3}
	want := []int{1, 2, 3, 4, 5, 6}
	got := clrs.SelectionSort(input)
	if !EqualIntSlices(got, want) {
		t.Fatalf("got:\t%v\nwant:\t%v\n", got, want)
	}
	input = []int{31, 41, 59, 26, 41, 58}
	want = []int{26, 31, 41, 41, 58, 59}
	got = clrs.SelectionSort(input)
	if !EqualIntSlices(got, want) {
		t.Fatalf("got:\t%v\nwant:\t%v\n", got, want)
	}
}

func EqualIntSlices(a, b []int) bool {
	for i, num := range a {
		if num != b[i] {
			return false
		}
	}
	return true
}
