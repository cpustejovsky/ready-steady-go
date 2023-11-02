package split_test

import (
	"github.com/cpustejovsky/ready-steady-go/split"
	"reflect"
	"testing"
)

var str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var want = []string{"ABC", "DEF", "GHI", "JKL", "MNO", "PQR", "STU", "VWX"}

func TestSplitString(t *testing.T) {
	got := split.String(str, 3)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf(" \ngot:\t%v\nwant:\t%v\n", got, want)
	}
}

func TestSplitStringUsingRunes(t *testing.T) {
	got := split.StringUsingRunes(str, 3)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf(" \ngot:\t%v\nwant:\t%v\n", got, want)
	}
}

func BenchmarkSplitString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		got := split.String(str, 3)
		if !reflect.DeepEqual(got, want) {
			b.Fatalf(" \ngot:\t%v\nwant:\t%v\n", got, want)
		}
	}
}
func BenchmarkSplitStringUsingRunes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		got := split.StringUsingRunes(str, 3)
		if !reflect.DeepEqual(got, want) {
			b.Fatalf(" \ngot:\t%v\nwant:\t%v\n", got, want)
		}
	}
}
