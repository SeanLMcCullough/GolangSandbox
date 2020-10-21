package route

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
)

func setup() *http.ServeMux {
	log, _ := test.NewNullLogger()
	mux := http.NewServeMux()
	ConfigureRoutes(mux, log)
	return mux
}

func TestGetHealth(t *testing.T) {
	mux := setup()
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal("request for GET /health failed")
	}
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("wrong status code. wanted %d, got %d", http.StatusOK, rec.Code)
	}
	expected := []byte("{\"Healthy\":true}\n")
	if response := rec.Body.Bytes(); !bytes.Equal(response, expected) {
		t.Errorf("response from GET /health was incorrect. wanted %q but got %q", expected, response)
	}
}
