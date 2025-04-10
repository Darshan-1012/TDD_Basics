package v2

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"
)

// SpyStore allows you to simulate a store and see how its used.
/* Second
Fetch returns response after a short delay.
// func (s *SpyStore) Fetch() string {
// 	time.Sleep(100 * time.Millisecond)
// 	return s.response
// }

// Cancel will record the call.
// func (s *SpyStore) Cancel() {
// 	s.cancelled = true
// }

// func (s *SpyStore) assertWasCancelled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Error("store was not told to cancel")
// 	}
// }

// func (s *SpyStore) assertWasNotCancelled() {
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Error("store was told to cancel")
// 	}
// }
*/
type SpyStore struct {
	response string
	// cancelled bool
	// t         *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)
	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(100 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

// WriteHeader implements http.ResponseWriter.
func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	panic("unimplemented")
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemnted")
}

func (s *SpyResponseWriter) WriterHeader(statusCode int) {
	s.written = true
}
