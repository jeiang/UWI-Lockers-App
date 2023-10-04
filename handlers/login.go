package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func loginHandler(w http.ResponseWriter, _ *http.Request) {
	log.Printf("loging??\n")
	fmt.Fprintf(w, "Hello, from login!")
}
