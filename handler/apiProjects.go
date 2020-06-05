package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tteeoo/go-website/limit"
	"log"
	"net/http"
)

// Put projects to appear on projects page here
var projects = []string{"rco", "sest", "go-website", "jschess", "claymore"}

var limiter = limit.NewIPRateLimiter(1, 1)
var send = []byte{}

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

// APIProjectHandler handles the /api/projects page
func APIProjectHandler(w http.ResponseWriter, r *http.Request) {

	// Rate limit
	// Don't use GitHub API if too fast
	// Uses last sent data
	limiter := limiter.GetLimiter(r.RemoteAddr)
	if !limiter.Allow() {
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, string(send))
		return
	}

	// Get all my repos
	gRepos, _, err := client.Repositories.List(context.Background(), "tteeoo", nil)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
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
	send, err = json.Marshal(repos)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json")
	fmt.Fprint(w, string(send))
}
