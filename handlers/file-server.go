package handlers

import (
	"net/http"
)

type FileServerWithFallback struct {
	HttpFileServer http.FileSystem
	Fallback       string
}

func (fswf *FileServerWithFallback) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serveIndex := ServeHTMLFile(fswf.Fallback, fswf.HttpFileServer)
	interceptor := Intercept(404, http.FileServer(fswf.HttpFileServer), serveIndex)
	interceptor.ServeHTTP(w, r)
}
