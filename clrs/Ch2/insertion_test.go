package Ch2_test

import (
	"github.com/cpustejovsky/ready-steady-go/clrs/Ch2"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	input := []int{5, 2, 4, 6, 1, 3}
	want := []int{1, 2, 3, 4, 5, 6}
	got := Ch2.InsertionSort(input)
	if !EqualIntSlices(want, got) {
		t.Fatalf("got:\t%v\nwant:\t%v\n", got, want)
	}
	input = []int{31, 41, 59, 26, 41, 58}
	want = []int{26, 31, 41, 41, 58, 59}
	got = Ch2.InsertionSort(input)
	if !EqualIntSlices(want, got) {
		t.Fatalf("got:\t%v\nwant:\t%v\n", got, want)
	}
}

func TestReverseInsertionSort(t *testing.T) {
	input := []int{5, 2, 4, 6, 1, 3}
	want := []int{6, 5, 4, 3, 2, 1}
	got := Ch2.ReverseInsertionSort(input)
	if !EqualIntSlices(want, got) {
		t.Fatalf("got:\t%v\nwant:\t%v\n", got, want)
	}
	input = []int{31, 41, 59, 26, 41, 58}
	want = []int{26, 31, 41, 41, 58, 59}
	want = []int{59, 58, 41, 41, 31, 26}
	got = Ch2.ReverseInsertionSort(input)
	if !EqualIntSlices(want, got) {
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
