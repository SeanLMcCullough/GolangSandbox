package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogGetRequest(t *testing.T) {
	handlerCalled := false
	h := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
		handlerCalled = true
	})
	r := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	RequestLoggerMiddleware(h).ServeHTTP(w, r)
	if !handlerCalled {
		t.Error("middleware must call handler")
	}
}
