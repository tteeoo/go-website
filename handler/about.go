package handler

import (
	"fmt"
	"net/http"
)

// AboutHandler handles the /about page
func AboutHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, aboutHTML)
}
