package download

import (
	"testing"
)

func TestHello(t *testing.T) {

	got := 8
	want := 8

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}
