package handler

import (
	"github.com/tteeoo/go-website/util"
	"html/template"
	"net/http"
	"strconv"
)

type errorPage struct {
	Errno int
	Text  string
}

// ErrorHandler handles errors by taking a status code and rendering a template with text.
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {

	util.Logger.Println("ERROR: " + util.GetRemoteAddr(r) + " " + strconv.Itoa(status))

	w.WriteHeader(status)

	// Fill in error template with the error number and text
	t, err := template.New("error").Parse(html["top"] + html["error"] + html["bottom"])
	if err != nil {
		// If this errors then bad stuff will likely happen
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
	}
}
