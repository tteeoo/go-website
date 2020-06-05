package handler

import (
	"fmt"
	"net/http"
)

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, projectsHTML)
}
