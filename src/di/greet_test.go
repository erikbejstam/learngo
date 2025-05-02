package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buf := bytes.Buffer{}
	Greet(&buf, "Erik")

	got := buf.String()
	want := "Hello, Erik"

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
