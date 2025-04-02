package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returns the fast", func(t *testing.T) {
		slowServer := makeDelayServer(20 * time.Millisecond)
		fastServer := makeDelayServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatal("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if server didn't respond", func(t *testing.T) {
		// serverA := makeDelayServer(11 * time.Second)
		// serverB := makeDelayServer(12 * time.Second)
		// Old line above

		server := makeDelayServer(25 * time.Millisecond)

		// defer serverA.Close()
		// defer serverB.Close()

		defer server.Close()

		// _, err := Racer(serverA.URL, serverB.URL)

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay) // time.Sleep(20* time.Milliseconds)
		w.WriteHeader(http.StatusOK)
	}))
}
