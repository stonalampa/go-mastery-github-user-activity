package api

import (
	"io"
	"net/http"
)

// * define the interface
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

// * API client struct
type GithubAPI struct {
	client  HTTPClient
	baseURL string
}

// * constructor with Dependency Injection
func NewGithubAPI(client HTTPClient) *GithubAPI {
	if client == nil {
		client = &http.Client{}
	}
	return &GithubAPI{
		client:  client,
		baseURL: "https://api.github.com",
	}
}

// * receive a username and return the user's activity
// * adding method to the GithubAPI struct
func (g *GithubAPI) FetchUserActivity(username string) ([]byte, error) {
	url := g.baseURL + "/users/" + username + "/events"

	resp, err := g.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
