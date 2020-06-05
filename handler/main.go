package handler

import (
	"context"
	"encoding/json"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"os"
)

var client *github.Client

var colors map[string]string

var indexHTML string
var errorHTML string
var aboutHTML string
var projectsHTML string

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
	aboutHTML = readHTML("about")
	errorHTML = readHTML("error")
	projectsHTML = readHTML("projects")

	// Read colors file
	b, err := ioutil.ReadFile("./static/colors.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(b, &colors)
	if err != nil {
		log.Fatal(err)
	}

	// Create GitHub client
	token, exists := os.LookupEnv("GH_WEB_TOKEN")
	if !exists {
		log.Fatal("GH_WEB_TOKEN not in environment")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client = github.NewClient(tc)
}
