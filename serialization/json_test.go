package serialization

import (
	"net/http/httptest"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	rec := httptest.NewRecorder()
	payload := struct {
		a string
		b int
		c bool
	}{
		a: "test",
		b: 100,
		c: true,
	}
	WriteJSON(rec, payload)
	expectedContentType := "application/json"
	if contentType := rec.Result().Header.Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Content-Type header not set correctly. wanted %q but got %q", expectedContentType, contentType)
	}
	// expected := []byte("{\"Healthy\":true}\n")
	// if response := rec.Body.Bytes(); !bytes.Equal(response, expected) {
	// 	t.Errorf("payload malformed. wanted %q but got %q", expected, response)
	// }
}
