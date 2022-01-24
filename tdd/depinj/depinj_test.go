package depinj_test

import (
	"bytes"
	"testing"

	"github.com/cpustejovsky/ready-steady-go/tdd/depinj"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	depinj.Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
