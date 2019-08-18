package hello

import (
	"encoding/json"
	"net/http"
)

func (s *server) handleRoot() http.HandlerFunc {
	type endpoint struct {
		Method string `json:"method,omitempty"`
		Path   string `json:"path,omitempty"`
	}
	type response struct {
		Endpoints []endpoint `json:"endpoints,omitempty"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		endpoints := []endpoint{
			{
				Method: "GET",
				Path:   "/hello",
			},
		}
		resp := response{endpoints}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			code := http.StatusInternalServerError
			http.Error(w, "unable to encode response", code)
		}
	}
}

func (s *server) handleHello() http.HandlerFunc {
	type response struct {
		Msg string `json:"msg,omitempty"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		resp := response{"Hello World!"}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			code := http.StatusInternalServerError
			http.Error(w, "unable to encode response", code)
		}
	}
}
