package unit

import (
	"testing"

	outputMapper "main/output-mapper"
)

func TestMapUserActivity(t *testing.T) {
	// * create mock data to test the mapper
	mockData := []byte(`[
		{
			"type": "PushEvent",
			"repo": {
				"name": "test-repo"
			}
		},
		{
			"type": "PushEvent",
			"repo": {
				"name": "test-repo-2"
			}
		}
	]`)

	// * call the mapper with the mock data
	mappedEvents, err := outputMapper.MapUserActivity(mockData)
	if err != nil {
		t.Errorf("Error mapping user activity: %v", err)
	}

	// *assert the mapped events
	if len(mappedEvents) != 2 {
		t.Errorf("Expected 2 mapped event, got %d", len(mappedEvents))
	}

	if mappedEvents[0] != "Event: PushEvent, Repository: test-repo" {
		t.Errorf("Expected 'Event: PushEvent, Repository: test-repo', got %s", mappedEvents[0])
	}

	if mappedEvents[1] != "Event: PushEvent, Repository: test-repo-2" {
		t.Errorf("Expected 'Event: PushEvent, Repository: test-repo-2', got %s", mappedEvents[1])
	}
}
