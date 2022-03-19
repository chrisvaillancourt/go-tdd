package hello

import "testing"

func TestHello(t *testing.T) {
	// passing in `t *testing.T` tells the test code to fail when we need it to
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// testing.TB allows us to call helper fns from a test or benchmark
		// `t.Helper()`` tells the test suite this fn is a helper. That way, test
		// failures report the line number where it's called instead of in the helper
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)

	})
	t.Run("say 'Hello, World' when empty string supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)

	})
	t.Run("In spanish", func(t *testing.T) {
		got := Hello("Andres", "Spanish")
		want := "Hola, Andres"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In french", func(t *testing.T) {
		got := Hello("Penelope", "French")
		want := "Bonjour, Penelope"
		assertCorrectMessage(t, got, want)
	})

}
