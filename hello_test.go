package main

import "testing"

// Single test example

// func TestHello(t *testing.T) {
// 	got := Hello("Shaun")
// 	want := "Hello, Shaun"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// Test with sub tests

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Shaun", "French") 
		want := "Bonjour, Shaun"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Italian", func(t *testing.T) {
		got := Hello("Liam", "Italian") 
		want := "Ciao, Liam"
		assertCorrectMessage(t, got, want)
	})
}

// Test helper function 

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // tell test suite this is a helper so use function call for text fail line number
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// Note: testing.TB is an interface testing.T and testing.B support
