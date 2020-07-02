package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/tteeoo/go-website/util"
)

var token string

var colors map[string]string

var indexHTML string
var errorHTML string
var aboutHTML string
var projectsHTML string

// readHTML reads an html file and returns the contents as a string
// log.Fatal on error
func readHTML(name string) string {

	file := "./template/" + name + ".html"

	b, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	util.Logger.Println("Read HTML file: " + file)

	return string(b)
}

func init() {

	// Read HTML into memory for faster requests
	indexHTML = readHTML("index")
	aboutHTML = readHTML("about")
	errorHTML = readHTML("error")
	projectsHTML = readHTML("projects")

	// Read colors file
	b, err := ioutil.ReadFile("./static/colors.json")
	if err != nil {
		util.Logger.Fatal(err)
	}

	err = json.Unmarshal(b, &colors)
	if err != nil {
		util.Logger.Fatal(err)
	}

	// Create GitHub client
	var exists bool
	token, exists = os.LookupEnv("GH_WEB_TOKEN")
	if !exists {
		util.Logger.Fatal("GH_WEB_TOKEN not in environment")
	}
}
