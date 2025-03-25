package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	StartA := time.Now()
	http.Get(a)
	aDuration := time.Since(StartA)

	StartB := time.Now()
	http.Get(b)
	bDuration := time.Since(StartB)

	if aDuration < bDuration {
		return a
	}
	return b
}
