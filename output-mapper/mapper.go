package outputMapper

import (
	"encoding/json"
	"fmt"
)

type Event struct {
	Id string `json:"id"`
	CreatedAt string `json:"created_at"`
}

func MapUserActivity(response []byte) ([]string, error) {
    var events []map[string]interface{}
    if err := json.Unmarshal(response, &events); err != nil {
        return nil, err
    }

    var mappedEvents []string
    for _, event := range events {
        eventType := event["type"].(string)
        repoName := event["repo"].(map[string]interface{})["name"].(string)

        mappedEvent := fmt.Sprintf("Event: %s, Repository: %s", eventType, repoName)
        mappedEvents = append(mappedEvents, mappedEvent)
    }

    return mappedEvents, nil
}
