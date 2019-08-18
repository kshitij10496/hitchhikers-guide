package hello

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	logger *log.Logger
}

func NewServer(r *mux.Router) *server {
	s := &server{
		router: r,
		logger: log.New(),
	}
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	s.logger.SetFormatter(formatter)
	s.routes()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
