package duplicate

import (
	"log"
)

var visited = make(map[string]bool)

type SimpleDuplicator struct {
}

func (d *SimpleDuplicator) Duplicate(url string) bool {
	if visited[url] {
		log.Printf("Duplicate: duplicated url %s", url)
		return true
	}

	visited[url] = true
	return false
}
