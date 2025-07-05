package main

import "testing"

func TestHello(t *testing.T) {
	got := 2 + 2
	want := 4
	if got != want {
		t.Errorf("expected %v, got %v", want, got)
	}
}
