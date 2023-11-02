package leetcode

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	want := "fl"
	got := longestCommonPrefix([]string{"flower", "flow", "flight"})
	if got != want {
		t.Fatalf("got %s, wanted %s", got, want)
	}
	want = ""
	got = longestCommonPrefix([]string{"dog", "racecar", "car"})
	if got != want {
		t.Fatalf("got %s, wanted %s", got, want)
	}
	want = ""
	got = longestCommonPrefix([]string{"reflower", "flow", "flight"})
	if got != want {
		t.Fatalf("got %s, wanted %s", got, want)
	}
	want = "aa"
	got = longestCommonPrefix([]string{"aaa", "aa", "aaa"})
	if got != want {
		t.Fatalf("got %s, wanted %s", got, want)
	}
	want = ""
	got = longestCommonPrefix([]string{"c", "acc", "ccc"})
	if got != want {
		t.Fatalf("got %s, wanted %s", got, want)
	}
}
