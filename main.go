package main

import (
	"log"
	"net/http"
	"github.com/tteeoo/go-website/handler"
)

func main() {

	// Handle routes
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/static/", handler.StaticHandler)

	// Start the server
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}
