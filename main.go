package main

import (
	"github.com/tteeoo/go-website/handler"
	"github.com/tteeoo/go-website/limit"
	"github.com/tteeoo/go-website/util"
	"net/http"
	"os"
)

var addr string

var limiter = limit.NewIPRateLimiter(1, 5)

func init() {

	// Get addr if set
	addr = os.Getenv("WEB_ADDR")
	if len(addr) == 0 {
		addr = "127.0.0.1:8000"
	}
}

func main() {

	// Setup logger
	defer util.LogFile.Close()

	// Handle routes
	http.HandleFunc("/", rateLimit(handler.IndexHandler))
	http.HandleFunc("/about", rateLimit(handler.AboutHandler))
	http.HandleFunc("/projects", rateLimit(handler.ProjectsHandler))
	http.HandleFunc("/robots.txt", rateLimit(handler.RootHandler))
	http.HandleFunc("/sitemap.xml", rateLimit(handler.RootHandler))

	http.HandleFunc("/static/", handler.StaticHandler)
	http.HandleFunc("/api/projects", handler.APIProjectHandler)

	// Start the server
	util.Logger.Println("Attempting to listen on http://" + addr)
	util.Logger.Fatal(http.ListenAndServe(addr, nil))
}

func rateLimit(handle func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		limiter := limiter.GetLimiter(util.GetRemoteAddr(r))
		if !limiter.Allow() {
			handler.ErrorHandler(w, r, http.StatusTooManyRequests)
			return
		}

		util.Logger.Println("HIT: " + util.GetRemoteAddr(r) + " " + r.RequestURI)
		handle(w, r)
	}
}
