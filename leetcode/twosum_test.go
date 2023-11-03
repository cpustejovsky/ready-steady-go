package leetcode_test

import (
	"github.com/cpustejovsky/ready-steady-go/leetcode"
	"testing"
)

func TestTwoSum(t *testing.T) {
	want := []int{0, 1}
	t.Run("Brute Force", func(t *testing.T) {
		got := leetcode.TwoSum([]int{2, 4, 11, 3}, 6)
		if !(got[0] == want[0] || got[0] == want[1]) {
			t.Fatalf("got:\t%v\nwant:\t%v\n", got, want)
		}
		if !(got[1] == want[0] || got[1] == want[1]) {
			t.Fatalf("got:\t%v\nwant:\t%v\n", got, want)
		}
	})
	t.Run("Two-Pass Hash Table", func(t *testing.T) {
		got := leetcode.TwoSumsTwoPassHashTable([]int{2, 4, 11, 3}, 6)
		if !(got[0] == want[0] || got[0] == want[1]) {
			t.Fatalf("got:\t%v\nwant:\t%v\n", got, want)
		}
		if !(got[1] == want[0] || got[1] == want[1]) {
			t.Fatalf("got:\t%v\nwant:\t%v\n", got, want)
		}
	})
	t.Run("One Pass Hash Table", func(t *testing.T) {
		got := leetcode.TwoSumsOnePassHashTable([]int{2, 4, 11, 3}, 6)
		if !(got[0] == want[0] || got[0] == want[1]) {
			t.Fatalf("got:\t%v\nwant:\t%v\n", got, want)
		}
		if !(got[1] == want[0] || got[1] == want[1]) {
			t.Fatalf("got:\t%v\nwant:\t%v\n", got, want)
		}
	})
}

func BenchmarkTwoSums(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leetcode.TwoSum([]int{2, 4, 11, 3}, 6)
	}
}

func BenchmarkTwoSumsTwoPassHashTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leetcode.TwoSumsTwoPassHashTable([]int{2, 4, 11, 3}, 6)
	}
}

func BenchmarkTwoSumsOnePassHashTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leetcode.TwoSumsOnePassHashTable([]int{2, 4, 11, 3}, 6)
	}
}
