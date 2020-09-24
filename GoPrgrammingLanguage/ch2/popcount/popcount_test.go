package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	for i := 0; i < 256; i++ {
		got := TablePopCount(uint64(i))
		want := PopCount(uint64(i))

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
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
