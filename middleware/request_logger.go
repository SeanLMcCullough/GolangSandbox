package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// RequestLoggerMiddleware logs http requests
func RequestLoggerMiddleware(log *logrus.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(logrus.Fields{
			"User-Agent": r.Header.Get("User-Agent"),
		}).Infof("%s %s\n", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
