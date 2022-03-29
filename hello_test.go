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
		got := Hello("James")
		want := "Hello James!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Greet with 'Hello Go!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello Go!"

		assertCorrectMessage(t, got, want)
	})
}
