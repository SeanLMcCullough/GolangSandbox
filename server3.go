package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/SeanLMcCullough/GoMicro/middleware"
	log "github.com/sirupsen/logrus"
)

const requestStartRequestKey = "requestStart"
const httpPort = 8000

// Health represents the current state of the running service, and its readiness to accept requests.
type Health struct {
	Healthy bool
}

func init() {
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", middleware.RequestLoggerMiddleware(http.HandlerFunc(handler)))
	mux.Handle("/health", middleware.RequestLoggerMiddleware(http.HandlerFunc(healthcheckHandler)))
	log.Infof("listening on %d", httpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Warning(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	health := Health{Healthy: true}
	writeJSON(w, health)
}

func writeJSON(writer http.ResponseWriter, payload interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(payload)
}
