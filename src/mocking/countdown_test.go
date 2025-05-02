package mocking

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buf := bytes.Buffer{}
	s := SpySleeper{}

	Countdown(&buf, &s)

	got := buf.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}

	if s.Calls != 3 {
		t.Errorf("Wanted 3 calls to SpySleeper, got %d", s.Calls)
	}
}
