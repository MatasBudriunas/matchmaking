package routes

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/", http.FileServer(http.Dir("./static")))
}
