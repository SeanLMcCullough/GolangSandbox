package main

import (
	"fmt"
	"net/http"

	"github.com/SeanLMcCullough/GoMicro/logging"

	"github.com/SeanLMcCullough/GoMicro/route"
)

const requestStartRequestKey = "requestStart"
const httpPort = 8000

func main() {
	log := logging.NewLogger()
	mux := http.NewServeMux()
	route.ConfigureRoutes(mux, log)
	log.Infof("listening on %d", httpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux))
}
