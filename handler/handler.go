package handler

import (
	"io/ioutil"

	"github.com/tteeoo/go-website/util"
)

var indexHTML string
var errorHTML string
var aboutHTML string
var projectsHTML string

// readHTML reads an html file and returns the contents as a string.
func readHTML(name string) string {

	file := "./template/" + name + ".html"

	b, err := ioutil.ReadFile(file)

	if err != nil {
		util.Logger.Fatalln(err)
	}

	util.Logger.Println("Read HTML file: " + file)

	return string(b)
}

func init() {

	// Read HTML into memory for faster requests
	indexHTML = readHTML("index")
	errorHTML = readHTML("error")
}
