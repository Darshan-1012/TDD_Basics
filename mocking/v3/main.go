package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const finalWord = "Go!"

// const countdownStart = 30

// func Countdown(out io.Writer, sleeper Sleeper) {
// 	for i := countdownStart; i > 0; i-- {
// 		fmt.Fprintln(out, i)
// 		sleeper.Sleep()
// 	}
// 	fmt.Fprint(out, finalWord)
// }

// The new count down for 1.23+ iterator usage
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := range countDownFrom(3) {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

// This is the iterator for the above countdownfrom

func countDownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
