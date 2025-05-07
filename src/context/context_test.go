package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Errorf("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Errorf("store was told to cancel")
	}
}

func TestServer(t *testing.T) {
	t.Run("return data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, cancelled: false, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil) // make a mock request
		response := httptest.NewRecorder()                       // make a mock response

		svr.ServeHTTP(response, request) // execute the two above

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		store.assertWasNotCancelled()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, cancelled: false, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		cancellingCtx, cancel := context.WithCancel(request.Context()) // wrap original context in a cancel context.
		time.AfterFunc(5*time.Millisecond, cancel)                     // wait and the runs cancel() (in go routine)
		request = request.WithContext(cancellingCtx)                   //create the new request with cancel ctx.

		svr.ServeHTTP(response, request)

		store.assertWasCancelled()
	})
}
