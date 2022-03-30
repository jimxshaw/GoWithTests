package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Greeting users", func(t *testing.T) {
		got := Hello("James", "English")
		want := "Hello James"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Greet with 'Hello Go' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello Go"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Marl", "Spanish")
		want := "Hola Marl"
		assertCorrectMessage(t, got, want)
	})
}
