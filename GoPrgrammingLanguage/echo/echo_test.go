// TODO: Implement Benchmark Testing for Echo Programming
package main

import (
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1()
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3()
	}
}

func BenchmarkEchoCustom(b *testing.B) {
	// os.Args = []string{"echo.go", "Hello", "World", "Test"}
	for i := 0; i < b.N; i++ {
		customEcho()
	}
}
