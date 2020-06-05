package handler

import (
	"html/template"
	"net/http"
	"github.com/tteeoo/go-website/util"
	"strconv"
)

type errorPage struct {
	Errno int
	Text  string
}

// ErrorHandler handles errors by taking a status code and rendering a template with text
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {

	util.Logger.Println(r.RemoteAddr + " got HTTP Error: " + strconv.Itoa(status))

	w.WriteHeader(status)

	// Fill in error template with the error nubmer and text
	t, err := template.New("error").Parse(errorHTML)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	ep := errorPage{
		Errno: status,
		Text:  http.StatusText(status),
	}

	err = t.Execute(w, ep)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}
