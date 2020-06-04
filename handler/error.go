package handler

import (
	"log"
	"net/http"
	"html/template"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {

	w.WriteHeader(status)

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
