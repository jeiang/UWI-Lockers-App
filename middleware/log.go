package middleware

import (
	"net/http"

	"github.com/charmbracelet/log"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug("received request", "method", r.Method, "path", r.URL.Path, "addr", r.RemoteAddr)

		next.ServeHTTP(w, r)
	})
}
