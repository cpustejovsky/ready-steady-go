package split_test

import (
	"github.com/cpustejovsky/ready-steady-go/split"
	"reflect"
	"testing"
)

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var emojis = "😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀"

func TestSplitString(t *testing.T) {
	t.Run("Alphabet", func(t *testing.T) {
		var want = []string{"ABC", "DEF", "GHI", "JKL", "MNO", "PQR", "STU", "VWX", "YZ"}

		got := split.String(alphabet, 3)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf(" \ngot:\t%v\nwant:\t%v\n", got, want)
		}
	})
	t.Run("Emojis", func(t *testing.T) {
		want := []string{"😀😀😀", "😀😀😀", "😀😀😀", "😀😀😀", "😀😀😀", "😀😀😀", "😀😀😀", "😀😀😀", "😀😀😀", "😀😀😀", "😀😀😀"}
		got := split.String(emojis, 3)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf(" \ngot:\t%v\nwant:\t%v\n", got, want)
		}
	})
}

func BenchmarkSplitString(b *testing.B) {
	var want = []string{"ABC", "DEF", "GHI", "JKL", "MNO", "PQR", "STU", "VWX", "YZ"}
	for i := 0; i < b.N; i++ {
		got := split.String(alphabet, 3)
		if !reflect.DeepEqual(got, want) {
			b.Fatalf(" \ngot:\t%v\nwant:\t%v\n", got, want)
		}
	}
}
