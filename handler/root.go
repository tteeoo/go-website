package handler

import (
	"net/http"
)

// RootHandler handles the static files at the root
func RootHandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "static"+r.URL.Path)
}
