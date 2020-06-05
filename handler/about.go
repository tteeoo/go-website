package handler

import (
	"fmt"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, aboutHTML)
}
