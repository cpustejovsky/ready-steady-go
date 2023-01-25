package clrs_test

import (
	"github.com/cpustejovsky/ready-steady-go/clrs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	input := []int{5, 2, 4, 6, 1, 3}
	want := []int{1, 2, 3, 4, 5, 6}
	got := clrs.InsertionSort(input)
	assert.Equal(t, want, got)
	input = []int{31, 41, 59, 26, 41, 58}
	want = []int{26, 31, 41, 41, 58, 59}
	got = clrs.InsertionSort(input)
	assert.Equal(t, want, got)
}
