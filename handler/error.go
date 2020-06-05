package handler

import (
	"html/template"
	"log"
	"net/http"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {

	w.WriteHeader(status)

	// Fill in error template with the error nubmer
	t, err := template.New("error").Parse(errorHTML)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, status)
	if err != nil {
		log.Println(err)
		return
	}
}
