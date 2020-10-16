package handler

import (
	"fmt"
	"net/http"
	"github.com/tteeoo/go-website/util"
)

// IndexHandler handles the / page.
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	util.Logger.Println("HIT: " + util.GetRemoteAddr(r) + " " + r.RequestURI)

	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	fmt.Fprint(w, indexHTML)
}
