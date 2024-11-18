package api

import (
	"io"
	"net/http"
)

func FetchUserActivity(username string) ([]byte, error) {
	url := "https://api.github.com/users/" + username + "/events"

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return io.ReadAll(response.Body)
}
