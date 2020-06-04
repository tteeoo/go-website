package handler

import (
	"log"
	"io/ioutil"
)

var indexHTML string
var errorHTML string

// readHTML reads an html file and returns the contents as a string
// log.Fatal on error
func readHTML(name string) string {

	file := "./templates/" + name + ".html"

	b, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Read HTML file: " + file)

	return string(b)
}

func init() {

	// Read HTML into memory for faster requests
	indexHTML = readHTML("index")
	errorHTML = readHTML("error")
}

