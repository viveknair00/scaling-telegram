package utils

import (
	"testing"
)

func TestAdd(t *testing.T) {

	got := Add(5, 3)
	want := 8

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}

func TestSubtract(t *testing.T) {

	got := Subtract(5, 5)
	want := 0

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}
