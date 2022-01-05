package main

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
