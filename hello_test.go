package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Malik")
	want := "Hello Malik"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
