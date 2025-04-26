package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const StartCount = 3
const EndCount = "Go!"

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := StartCount; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, EndCount)
}
func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
