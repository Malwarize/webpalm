package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Malwarize/webpalm/v2/core"
	"github.com/Malwarize/webpalm/v2/shared"
	"github.com/Malwarize/webpalm/v2/webtree"
)

func TestCache_AddAndIsVisited(t *testing.T) {
	cache := core.Cache{Visited: make(map[string]bool)}
	url := "http://example.com"

	// Initially, the URL should not be marked as visited
	assert.False(t, cache.IsVisited(url), "URL should not be marked as visited initially")

	// Add the URL to the cache
	cache.AddVisited(url)

	// Now, the URL should be marked as visited
	assert.True(t, cache.IsVisited(url), "URL should be marked as visited after adding")
}

func TestCrawler(t *testing.T) {
	options := &shared.Options{
		URL:             "http://example.com",
		Level:           2,
		ExportFile:      "",
		RegexMap:        map[string]string{},
		StatusResponses: []int{},
		IncludedUrls:    []string{},
		Workers:         0,
		Delay:           0,
		UserAgent:       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 cko) Chrome/70.0.3538.77 Safari/537.36",
		TimeOut:         10,
		Proxy:           nil,
	}
	c := core.NewCrawler(options)

	page := webtree.NewPage()
	page.SetUrl("http://example.com")
	c.Fetch(page)
	links := c.ExtractLinks(page)
	assert.NotEmpty(t, links, "ExtractLinks should return a non-empty list of links")

	node := webtree.NewNode(page, nil, nil)
	c.ProcessANode(node)

	assert.True(
		t,
		c.IsSkipableUrl("http://example.com/test.mp3"),
		"IsSkipableUrl should return true for http://example.com/test.mp3",
	)
	assert.True(t, c.IsSkipablePage(page), "IsSkipablePage should return true for the page")
	assert.Equal(t, options.Level, c.Level, "NewCrawler Level should match options.Level")
}
