package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3

	sleep = "sleep"
	write = "write"
)

//

type Sleeper interface {
	Sleep()
}

//

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration) // want to be able to pass mock sleep function
}

func NewConfigurableSleeper(d time.Duration, f func(time.Duration)) *ConfigurableSleeper {
	c := ConfigurableSleeper{
		duration: d,
		sleep:    f,
	}

	return &c
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

//

func Countdown(writer io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(writer, "%d\n", i)
		s.Sleep()
	}
	fmt.Fprint(writer, finalWord)
}
