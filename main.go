package main

import (
	"learn-go/src/mocking"
	"os"
	"time"
)

func main() {
	s := *mocking.NewConfigurableSleeper(1*time.Second, time.Sleep)
	mocking.Countdown(os.Stdout, &s)
}
