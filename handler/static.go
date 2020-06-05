package handler

import (
	"net/http"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {

	// if r.URL.Path != "/static/" {
	// 	errorHandler(w, r, http.StatusNotFound)
	// 	return
	// }

	http.ServeFile(w, r, r.URL.Path[1:])
}
