package duplicate

import (
	"testing"
)

func TestRedisDuplicate_NotDuplicate(t *testing.T) {
	expected := false

	// TODO: Try to start up redis here using docker go test.
	duplicator := NewRedisDuplicator("127.0.0.1:6379", "")

	url := "https://www.github.com"

	duplicator.Client.FlushDB(ctx)

	actual := duplicator.Duplicate(url)

	if actual != expected {
		t.Errorf("got %v; expected: %v", actual, expected)
	}
}

func TestRedisDuplicate_Duplicate(t *testing.T) {
	expected := true

	// TODO: Try to start up redis here using docker go test.
	duplicator := NewRedisDuplicator("127.0.0.1:6379", "")

	url := "https://www.github.com"

	duplicator.Client.FlushDB(ctx)

	duplicator.Duplicate(url)
	actual := duplicator.Duplicate(url)

	if actual != expected {
		t.Errorf("got %v; expected: %v", actual, expected)
	}
}
