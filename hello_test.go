package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		result := Hello("Matheus")
		expected := "Hello Matheus."

		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("with a empty name say Hello World", func(t *testing.T) {
		result := Hello("")
		expected := "Hello World."

		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
