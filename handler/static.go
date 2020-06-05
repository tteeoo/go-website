package handler

import (
	"net/http"
)

// StaticHandler handles the /static/ files
func StaticHandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, r.URL.Path[1:])
}
