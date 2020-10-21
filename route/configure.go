package route

import (
	"net/http"

	"github.com/SeanLMcCullough/GoMicro/middleware"
	"github.com/sirupsen/logrus"
)

// ConfigureRoutes configures the routes for the provided mux
func ConfigureRoutes(mux *http.ServeMux, log *logrus.Logger) {
	mux.Handle("/", middleware.RequestLoggerMiddleware(log, getRootRoute(log)))
	mux.Handle("/health", middleware.RequestLoggerMiddleware(log, getHealthRoute()))
}
