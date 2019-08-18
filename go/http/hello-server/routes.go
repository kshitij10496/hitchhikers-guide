package hello

import (
	"net/http"
)

func (s *server) routes() {
	s.router.HandleFunc("/", s.withLogging(s.handleRoot()))
	s.router.HandleFunc("/hello", s.withLogging(s.handleHello())).Methods(http.MethodGet)
}
