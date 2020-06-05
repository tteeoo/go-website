package main

// TODO:
// - Get the styles just right
// - Create Dockerfile

import (
	"github.com/tteeoo/go-website/handler"
	"github.com/tteeoo/go-website/limit"
	"github.com/tteeoo/go-website/util"
	"net/http"
)

const addr = "127.0.0.1:8000"

var limiter = limit.NewIPRateLimiter(1, 5)

func main() {

	// Setup logger
	defer util.LogFile.Close()
	logger := util.Logger

	// Handle routes
	http.HandleFunc("/", rateLimit(handler.IndexHandler))
	http.HandleFunc("/about", rateLimit(handler.AboutHandler))
	http.HandleFunc("/projects", rateLimit(handler.ProjectsHandler))

	http.HandleFunc("/static/", handler.StaticHandler)
	http.HandleFunc("/api/projects", handler.APIProjectHandler)

	// Start the server
	logger.Println("Attempting to listen on http://" + addr)
	logger.Fatal(http.ListenAndServe(addr, nil))
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
