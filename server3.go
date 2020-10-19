package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/SeanLMcCullough/GoMicro/middleware"
	log "github.com/sirupsen/logrus"
)

const requestStartRequestKey = "requestStart"
const httpPort = 8000

func init() {
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	mux := http.NewServeMux()
	finalHandler := http.HandlerFunc(handler)
	mux.Handle("/", middleware.RequestLoggerMiddleware(finalHandler))
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
