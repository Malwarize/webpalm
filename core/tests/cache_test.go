package core

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Malwarize/webpalm/v2/core"
)

func TestCache(t *testing.T) {
	cache := core.Cache{Visited: make(map[string]bool)}
	url := "http://example.com"

	// Add the URL to the cache
	cache.AddVisited(url)

	// Ensure the URL is marked as visited
	assert.True(t, cache.IsVisited(url), "URL should be marked as visited before flushing")

	// Flush the cache
	cache.Flush()

	// After flushing, the URL should not be marked as visited
	assert.False(t, cache.IsVisited(url), "URL should not be marked as visited after flushing")
}
