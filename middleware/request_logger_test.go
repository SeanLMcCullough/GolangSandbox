package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
)

var requestTests = []struct {
	method   string
	URL      string
	expected string
}{
	{"GET", "/test", "GET /test\n"},
	{"POST", "/asdf", "POST /asdf\n"},
	{"PUT", "/cat", "PUT /cat\n"},
	{"DELETE", "/dog", "DELETE /dog\n"},
}

func TestLogGetRequest(t *testing.T) {
	for _, testCase := range requestTests {
		log, hook := test.NewNullLogger()
		handlerCalled := false
		h := http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
			handlerCalled = true
		})
		r := httptest.NewRequest(testCase.method, testCase.URL, nil)
		w := httptest.NewRecorder()
		RequestLoggerMiddleware(log, h).ServeHTTP(w, r)

		if !handlerCalled {
			t.Errorf("%s %s did not call handler", testCase.method, testCase.URL)
		}
		if calls := len(hook.Entries); calls != 1 {
			t.Errorf("%s %s called %d times, wanted 1", testCase.method, testCase.URL, calls)
		}
		if level := hook.LastEntry().Level; level != logrus.InfoLevel {
			t.Errorf("%s %s called logger with level %s, wanted %s", testCase.method, testCase.URL, level, logrus.InfoLevel)
		}
		if message := hook.LastEntry().Message; message != testCase.expected {
			t.Errorf("%s %s = %s, wanted %s", testCase.method, testCase.URL, message, testCase.expected)
		}
	}
}
