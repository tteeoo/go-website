package handler

import (
	"fmt"
	"net/http"
)

// ProjectsHandler handles the /projects page
func ProjectsHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, projectsHTML)
}
