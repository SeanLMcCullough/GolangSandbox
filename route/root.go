package route

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func getRootRoute(log *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	})
}
