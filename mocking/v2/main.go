package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const StartCount = 3
const EndCount = "Go!"

func Countdown(out io.Writer) {
	for i := StartCount; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, EndCount)
}
func main() {
	Countdown(os.Stdout)
}
