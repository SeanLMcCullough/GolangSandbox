package serialization

import (
	"encoding/json"
	"net/http"
)

// WriteJSON writes the payload to the writer as JSON
func WriteJSON(writer http.ResponseWriter, payload interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(payload)
}
