package main

import (
	"os"
	"time"

	"example.com/tdd-exercises/mocking"
)

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
