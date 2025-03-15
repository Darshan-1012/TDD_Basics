package new

import (
	"testing"
)

func TestMain(t *testing.T) {
	// t.Run("The 'Hello Darshan' is executed", func(t *testing.T) {
	// 	got := hello("Darshan", "")
	// 	want := "Hello Darshan"
	// 	assertMsg(t, got, want)
	// })

	t.Run("The 'Hello World' is executed", func(t *testing.T) {
		got := hello("", "")
		want := "Hello World"
		assertMsg(t, got, want)
	})

	t.Run("hola in spain", func(t *testing.T) {
		got := hello("Darshan", "Spain")
		want := "Hola Darshan"
		assertMsg(t, got, want)
	})

	t.Run("bonjour in French", func(t *testing.T) {
		got := hello("Darshan", "French")
		want := "Bonjour Darshan"
		assertMsg(t, got, want)
	})
}

func assertMsg(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got is %q and want is %q", got, want)
	}
}
