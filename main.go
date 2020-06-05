package main

// TODO:
// - Add content
// - Get the styles just right
// - Create Dockerfile

import (
	"github.com/tteeoo/go-website/handler"
	"github.com/tteeoo/go-website/limit"
	"log"
	"net/http"
)

var limiter = limit.NewIPRateLimiter(1, 5)

func main() {

	// Handle routes
	http.HandleFunc("/", rateLimit(handler.IndexHandler))
	http.HandleFunc("/about", rateLimit(handler.AboutHandler))
	http.HandleFunc("/projects", rateLimit(handler.ProjectsHandler))

	http.HandleFunc("/static/", handler.StaticHandler)
	http.HandleFunc("/api/projects", handler.APIProjectHandler)

	// Start the server
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}

func rateLimit(handle func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		limiter := limiter.GetLimiter(r.RemoteAddr)
		if !limiter.Allow() {
			handler.ErrorHandler(w, r, http.StatusTooManyRequests)
			return
		}

		handle(w, r)
	}
}
