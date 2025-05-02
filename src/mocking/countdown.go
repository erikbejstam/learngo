package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(writer io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(writer, "%d\n", i)
		s.Sleep()
	}
	fmt.Fprint(writer, finalWord)
}
