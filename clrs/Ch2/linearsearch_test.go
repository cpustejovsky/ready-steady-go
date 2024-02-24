package Ch2_test

import (
	"github.com/cpustejovsky/ready-steady-go/clrs/Ch2"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 42}
	searchNum := 42
	want := 7
	got := Ch2.LinearSearch(arr, searchNum)
	if want != *got {
		t.Fatalf("wanted %v; got %v", want, *got)
	}

}
