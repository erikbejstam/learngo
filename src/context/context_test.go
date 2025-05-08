package context

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Store Spy

type SpyStore struct {
	data string
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	// simulate a slow process of fetching a bunch of data. think of this as doing the equivalent of "go get this and this data from this db".
	go func() {
		var result string
		for _, c := range s.data {
			select {
			case <-ctx.Done():
				fmt.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case result := <-data:
		return result, nil
	}
}

// ResponseWriter spy

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

// Tests

func TestServer(t *testing.T) {
	t.Run("get data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{data: data}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil) // get request, to target /, empty body
		response := httptest.NewRecorder()                       // just a test version of a ResponseWriter, I think

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got: %s, want: %s", response.Body.String(), data)
		}
	})

	t.Run("tells store to cancel if request cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{data: data}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)       // get request, to target /, empty body
		cancellingCtx, cancel := context.WithCancel(request.Context()) // unwieldy syntax for assigning a cancelctx to a request.
		time.AfterFunc(5*time.Millisecond, cancel)                     // But you have to take the (unmutable) ctx from the req, call the WithCancel function to get cancelCtx and cancel
		request = request.WithContext(cancellingCtx)                   // And then you take that context in this function, which returns the req you want.

		response := &SpyResponseWriter{false}

		server.ServeHTTP(response, request)

		if response.written {
			t.Errorf("a response was written")
		}
	})
}
