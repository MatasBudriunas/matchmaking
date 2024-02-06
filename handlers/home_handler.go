package handlers

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		NotFoundHandler(w)
		return
	}
	fmt.Fprintf(w, "Welcome to the Matchmaking Service")
}
