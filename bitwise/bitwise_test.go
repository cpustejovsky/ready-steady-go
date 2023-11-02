package bitwise_test

import (
	"github.com/cpustejovsky/ready-steady-go/bitwise"
	"testing"
)

func BenchmarkIsEvenBitwise(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitwise.IsEvenBitwise(i)
	}
}

func BenchmarkIsOddBitwise(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitwise.IsOddBitwise(i)
	}
}

func BenchmarkIsEvenModulo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitwise.IsEvenModulo(i)
	}
}

func BenchmarkIsOddModulo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitwise.IsOddModulo(i)
	}
}

func BenchmarkProcstr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		want := "!elpoeP olleH"
		got := bitwise.Procstr("HELLO PEOPLE!", bitwise.LOWER|bitwise.REV|bitwise.CAP)
		if got != want {
			b.Fatalf("got %s, wanted %s\n", got, want)
		}
	}
}
