package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(cxt context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			fmt.Printf("Got an error when fetching: %s", err.Error())
			return // todo: log error somehow
		}

		fmt.Fprint(w, data)
	}
}
