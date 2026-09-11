package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Matheus")
	expected := "Hello Matheus."

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
