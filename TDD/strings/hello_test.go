package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Charles", "")
		want := "Hello, Charles"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hola, Charles' when Spanish is provided as second parameter", func(t *testing.T) {
		got := Hello("Charles", "Spanish")
		want := "Hola, Charles"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Bonjour, Charles' when French is provided as second parameter", func(t *testing.T) {
		got := Hello("Charles", "French")
		want := "Bonjour, Charles"
		assertCorrectMessage(t, got, want)
	})

}
