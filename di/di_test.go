package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Darshan")

	got := buffer.String()
	want := "Hello Darshan"

	if got != want {
		t.Errorf("%q is got %q is want", got, want)
	}
}
