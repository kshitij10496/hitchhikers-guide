package hello

import (
	"net/http"
)

func (s *server) routes() {
	s.router.HandleFunc("/hello", s.handleHello()).Methods(http.MethodGet)
}
