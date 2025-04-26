package main

import (
	"bytes"
	"testing"
)

func TestCount(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("%q is got %q is want", got, want)
	}
	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls of sleeper, 3 is want %d is got", spySleeper.Calls)
	}
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
