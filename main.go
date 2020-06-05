package main

import (
	"github.com/tteeoo/go-website/handler"
	"log"
	"net/http"
)

func main() {

	// Handle routes
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/about", handler.AboutHandler)
	http.HandleFunc("/static/", handler.StaticHandler)
	http.HandleFunc("/projects", handler.ProjectsHandler)
	http.HandleFunc("/api/projects", handler.ApiProjectHandler)

	// Start the server
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}
