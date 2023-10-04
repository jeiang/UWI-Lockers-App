package handlers

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("received a request", "ipaddr", r.RemoteAddr)
	fmt.Fprintf(w, "Hello, from login!")
}
