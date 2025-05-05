package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("proper order", func(t *testing.T) {
		s := &SpyCountdownOperations{}

		Countdown(s, s)

		wantCalls := []string{write, sleep, write, sleep, write, sleep, write}

		if !reflect.DeepEqual(s.Calls, wantCalls) {
			t.Errorf("got: %v, want: %v", s.Calls, wantCalls)
		}
	})
	t.Run("output", func(t *testing.T) {
		s := &SpyCountdownOperations{}
		buf := &bytes.Buffer{}

		Countdown(buf, s)

		got := buf.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
