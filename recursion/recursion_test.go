package recursion_test

import (
	"github.com/cpustejovsky/ready-steady-go/recursion"
	"testing"
)

func TestSumUpRecursive(t *testing.T) {
	got := recursion.SumUpRecursive(10)
	want := 55
	if got != want {
		t.Fatalf("got %d, want %d\n", got, want)
	}
}

func TestSumUp(t *testing.T) {
	got := recursion.SumUp(10)
	want := 55
	if got != want {
		t.Fatalf("got %d, want %d\n", got, want)
	}
}

var fibMap = map[int]int{
	0:  0,
	1:  1,
	2:  1,
	3:  2,
	4:  3,
	5:  5,
	6:  8,
	7:  13,
	8:  21,
	9:  34,
	10: 55,
	11: 89,
	12: 144,
	13: 233,
	14: 377,
	15: 610,
	16: 987,
	17: 1597,
	18: 2584,
	19: 4181,
}

func TestFibonacci(t *testing.T) {
	for k, want := range fibMap {
		got := recursion.Fibonacci(k)
		if got != want {
			t.Fatalf("got %d, want %d\n", got, want)
		}
	}
}

func BenchmarkSumUpRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		got := recursion.SumUpRecursive(10)
		want := 55
		if got != want {
			b.Fatalf("got %d, want %d\n", got, want)
		}
	}

}

func BenchmarkSumUp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		got := recursion.SumUp(10)
		want := 55
		if got != want {
			b.Fatalf("got %d, want %d\n", got, want)
		}
	}
}
