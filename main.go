package main

import (
	"learn-go/src/mocking"
	"os"
)

func main() {
	s := mocking.DefaultSleeper{}
	mocking.Countdown(os.Stdout, &s)
}
