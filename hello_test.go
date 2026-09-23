package main

import "testing"

func TestHello(t *testing.T) {
	verifyCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		result := Hello("Matheus", "en")
		expected := "Hello, Matheus."
		verifyCorrectMessage(t, result, expected)
	})

	t.Run("with a empty name say Hello World", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello, World."
		verifyCorrectMessage(t, result, expected)
	})

	t.Run("greeting in spanish", func(t *testing.T) {
		result := Hello("Matheus", "es")
		expected := "Hola, Matheus."
		verifyCorrectMessage(t, result, expected)
	})
}
