package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tteeoo/go-website/limit"
	"github.com/tteeoo/go-website/util"
)

const query = "{ \"query\": \"query { user(login: \\\"tteeoo\\\") { pinnedItems(first: 6, types: [REPOSITORY]) { totalCount edges { node { ... on Repository { name url primaryLanguage { name } description } } } } } }\" }"

var limiter = limit.NewIPRateLimiter(1, 1)
var send = []byte{}

type apiRepo struct {
	Name  string
	Color string
	URL   string
	Desc  string
	Lang  string
}

type ghGQLLang struct {
	Name string `json:"name"`
}

type ghQGLRepoData struct {
	Name            string    `json:"name"`
	URL             string    `json:"url"`
	Description     string    `json:"description"`
	PrimaryLanguage ghGQLLang `json:"primaryLanguage"`
}

type ghQGLRepo struct {
	Data ghQGLRepoData `json:"node"`
}

type ghGQLResponse struct {
	Repos []ghQGLRepo `json:"edges"`
}

// APIProjectHandler handles the /api/projects page
func APIProjectHandler(w http.ResponseWriter, r *http.Request) {

	// Rate limit
	// Don't use GitHub API if too fast
	// Uses last sent data
	limiter := limiter.GetLimiter(util.GetRemoteAddr(r))
	if !limiter.Allow() {
		util.Logger.Println("API SPAM: " + util.GetRemoteAddr(r))
		w.Header().Set("content-type", "application/json")
		fmt.Fprint(w, string(send))
		return
	}

	// Use GH qraphql API
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.github.com/graphql", strings.NewReader(query))
	if err != nil {
		util.Logger.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
	}
	req.Header.Add("Authorization", "bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		util.Logger.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
	}

	// Read and deserialize response
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		util.Logger.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
	}
	var data map[string]map[string]map[string]ghGQLResponse
	err = json.Unmarshal(b, &data)
	if err != nil {
		util.Logger.Println(string(b))
		util.Logger.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
	}

	// Construct API response
	var apiRepos []*apiRepo
	repos := data["data"]["user"]["pinnedItems"].Repos
	for _, repo := range repos {
		r := &apiRepo{
			Name: repo.Data.Name,
			Desc: repo.Data.Description,
			URL:  repo.Data.URL,
			Lang: repo.Data.PrimaryLanguage.Name,
		}
		color, exists := colors[repo.Data.PrimaryLanguage.Name]
		if !exists {
			r.Color = "background-color: #383838"
		} else {
			r.Color = "background-color: " + color
		}
		apiRepos = append(apiRepos, r)
	}

	// Send repos as JSON
	send, err = json.Marshal(apiRepos)
	if err != nil {
		util.Logger.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json")
	fmt.Fprint(w, string(send))
}
