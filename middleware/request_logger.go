package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// RequestLoggerMiddleware logs http requests
func RequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"User-Agent": r.Header.Get("User-Agent"),
		}).Infof("%s %s\n", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
