package main

import "testing"

func TestHello(t *testing.T)  {
	asserCorrectMessage := func (t *testing.T, got, want string)  {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		asserCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		asserCorrectMessage(t, got, want)
	})
}