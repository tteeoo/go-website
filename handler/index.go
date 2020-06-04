package handler

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	fmt.Fprint(w, indexHTML)
}
