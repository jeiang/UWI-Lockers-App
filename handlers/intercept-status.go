package handlers

import (
	"net/http"
)

type hookedResponseWriter struct {
	http.ResponseWriter
	codeToIntercept  int
	gotInterceptCode bool
}

func (hrw *hookedResponseWriter) WriteHeader(status int) {
	if status == http.StatusNotFound {
		// Don't actually write the header, just set a flag.
		hrw.gotInterceptCode = true
	} else {
		hrw.ResponseWriter.WriteHeader(status)
	}
}

func (hrw *hookedResponseWriter) Write(p []byte) (int, error) {
	if hrw.gotInterceptCode {
		// No-op, but pretend that we wrote len(p) bytes to the writer.
		return len(p), nil
	}

	return hrw.ResponseWriter.Write(p)
}

func Intercept(code int, handler, onIntercept http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hookedWriter := &hookedResponseWriter{
			ResponseWriter:  w,
			codeToIntercept: code,
		}
		handler.ServeHTTP(hookedWriter, r)
		if hookedWriter.gotInterceptCode {
			onIntercept.ServeHTTP(w, r)
		}
	})
}
