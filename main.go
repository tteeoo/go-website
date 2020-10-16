package main

import (
	"github.com/tteeoo/go-website/handler"
	"github.com/tteeoo/go-website/util"
	"net/http"
	"os"
)

func main() {

	// Get addr if set
	var addr string
	addr = os.Getenv("WEB_ADDR")
	if len(addr) == 0 {
		addr = "127.0.0.1:8000"
	}

	// Setup logger
	defer util.LogFile.Close()

	// Handle routes
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/robots.txt", handler.FileHandler)
	http.HandleFunc("/sitemap.xml", handler.FileHandler)
	http.HandleFunc("/static/", handler.FileHandler)

	// Start the server
	util.Logger.Println("Attempting to listen at http://" + addr)
	util.Logger.Fatal(http.ListenAndServe(addr, nil))
}
