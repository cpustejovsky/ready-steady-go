package why_test

import (
	"testing"

	"github.com/cpustejovsky/ready-steady-go/tdd/why"
)

func TestHello(t *testing.T) {
	want := "Hola, Charles"
	got := why.Hello("Charles", "es")
	if got != want {
		t.Errorf("got:\t%v, wanted:\t%v", want, got)
	}
}
