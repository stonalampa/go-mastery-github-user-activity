package unit

import (
	"bytes"
	"io"
	"main/api"
	"net/http"
	"testing"
)

type MockHTTPClient struct{}

func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	// * mock response
	mockResponse := `[{"type": "PushEvent", "repo": {"name": "test/repo"}}]`
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(mockResponse)),
	}, nil
}

func TestFetchUserActivity(t *testing.T) {
	// * create mock client and API instance
	mockClient := &MockHTTPClient{}
	githubAPI := api.NewGithubAPI(mockClient)

	// * call the function
	response, err := githubAPI.FetchUserActivity("testuser")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// * assert response
	expectedResponse := `[{"type": "PushEvent", "repo": {"name": "test/repo"}}]`
	if string(response) != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, string(response))
	}
}
