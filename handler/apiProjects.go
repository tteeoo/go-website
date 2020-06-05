package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Put projects to appear on projects page here
var projects = []string{"rco", "sest", "go-website", "jschess", "claymore"}

type repo struct {
	Name  string
	Color string
	URL   string
	Desc  string
	Lang  string
}

func goodProject(a string) bool {
	for _, b := range projects {
		if b == a {
			return true
		}
	}
	return false
}

// ApiApiProjectHandler handles the /api/projects page
func ApiProjectHandler(w http.ResponseWriter, r *http.Request) {

	// Get all my repos
	gRepos, _, err := client.Repositories.List(context.Background(), "tteeoo", nil)
	if err != nil {
		log.Println(err)
		errorHandler(w, r, http.StatusInternalServerError)
	}

	// Iterate repos
	var repos []repo
	for _, gR := range gRepos {

		// continue if in is not in projects
		if !goodProject(*gR.Name) {
			continue
		}

		re := repo{
			Name: *gR.Name,
			URL:  *gR.HTMLURL,
		}

		// Ensure there are no nil pointers before dereferencing
		// Set fallback color if the language isn't in colors or is nil
		if gR.Language != nil {
			re.Lang = *gR.Language

			color, exists := colors[*gR.Language]
			if !exists {
				re.Color = "background-color: #383838"
			} else {
				re.Color = "background-color: " + color
			}
		} else {
			re.Color = "background-color: #383838"
		}

		if gR.Description != nil {
			re.Desc = *gR.Description
		}

		repos = append(repos, re)
	}

	// Send repos as JSON
	send, err := json.Marshal(repos)
	if err != nil {
		log.Println(err)
		errorHandler(w, r, http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json")
	fmt.Fprint(w, string(send))
}
