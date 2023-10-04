package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterApiRoutes(r *mux.Router) {
	// TODO: register some routes
	// The API will be served under `/api/v1`.
	s := r.PathPrefix("/api/v1").Subrouter()
	s.PathPrefix("/").Methods(http.MethodGet).HandlerFunc(handleApiIndex)
}

func handleApiIndex(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello, from api!")
}
