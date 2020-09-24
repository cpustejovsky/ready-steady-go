package popcount

import (
	"testing"
)

var tests = []struct {
	x    uint64
	want int
}{
	{0, 0},
	{1, 1},
	{1 << 8, 1},
	{1<<8 + 1, 2},
	{1<<8 - 1, 8},
	{1<<64 - 1, 64},
}

func TestPopCount(t *testing.T) {
	for _, test := range tests {
		if got := PopCount(test.x); got != test.want {
			t.Errorf("BitShiftPopCount(%d) = %d, want %d", test.x, got, test.want)
		}
	}
}

func TestTablePopCount(t *testing.T) {
	for _, test := range tests {
		if got := TablePopCount(test.x); got != test.want {
			t.Errorf("TablePopCount(%d) = %d, want %d", test.x, got, test.want)
		}
	}
}

func TestPopCountBitWiseAnd(t *testing.T) {
	for _, test := range tests {
		if got := PopCountBitWiseAnd(test.x); got != test.want {
			t.Errorf("PopCountBitWiseAnd(%d) = %d, want %d", test.x, got, test.want)
		}
	}
}

func BenchmarkTablePopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TablePopCount(uint64(i))
	}
}

//loop is way slower
func BenchmarkTablePopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TablePopCountLoop(uint64(i))
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}
