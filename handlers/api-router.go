package handlers

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gorilla/mux"
)

func RegisterApiRoutes(r *mux.Router) {
	// TODO: register some routes
	// The API will be served under `/api/v1`.
	s := r.PathPrefix("/api/v1").Subrouter()
	s.Path("/").Methods(http.MethodGet).HandlerFunc(handleApiIndex)
	s.Path("/login").Methods(http.MethodPost).HandlerFunc(loginHandler)
}

func handleApiIndex(w http.ResponseWriter, r *http.Request) {
	log.Debug("received a request", "ipaddr", r.RemoteAddr)
	fmt.Fprintf(w, "Hello, from api!")
}
