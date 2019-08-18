package hello

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func (s *server) withLogging(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.WithTime(time.Now()).WithFields(log.Fields{
			"method":   r.Method,
			"endpoint": r.URL.Path,
		}).Infof("")
		h(w, r)
	}
}
