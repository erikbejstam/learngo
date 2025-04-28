package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMsg(t, got, want)
	})
	
	t.Run("Defaulting to hello world", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMsg(t, got, want)
	})

	t.Run("In spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMsg(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Maurice", "French")
		want := "Bonjour, Maurice"
		assertCorrectMsg(t, got, want)
	})
}

func assertCorrectMsg(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}