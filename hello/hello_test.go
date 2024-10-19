package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Bryce")
	want := "Hello, Bryce!"

	if (got != want) {
		t.Errorf("got %q want %q", got, want)
	}
}