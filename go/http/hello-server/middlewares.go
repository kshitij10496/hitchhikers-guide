package hello

import "net/http"

func withLogging(n http.Handler) http.Handler {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
