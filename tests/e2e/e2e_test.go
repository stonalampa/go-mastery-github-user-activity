package e2e

import (
	"testing"

	"main/api"
)

func TestFetchUserActivity(t *testing.T) {
	// * create a new GithubAPI instance
	githubAPI := api.NewGithubAPI(nil)

	// * call the FetchUserActivity method
	response, err := githubAPI.FetchUserActivity("stonalampa")

	// * assert the response is not nil
	if response == nil {
		t.Errorf("Expected a response, got nil")
	}

	// * assert the error is nil
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
