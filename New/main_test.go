package new

import (
	"testing"
)

func TestMain(t *testing.T) {
	got := hello("Darshan")
	want := "Hello Darshan"

	if got != want {
		t.Errorf("%q is got \n %q is want ", got, want)
	}
}
