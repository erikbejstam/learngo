package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			fmt.Printf("Error occurred when fetching data from Store: %s", err.Error())
			return
		}
		fmt.Fprint(w, data)
	}
}
