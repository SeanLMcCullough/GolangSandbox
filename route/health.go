package route

import (
	"net/http"

	"github.com/SeanLMcCullough/GoMicro/serialization"
)

// Health represents the current state of the running service, and its readiness to accept requests.
type Health struct {
	Healthy bool
}

func getHealthRoute() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		health := Health{Healthy: true}
		serialization.WriteJSON(w, health)
	})
}
