package tests

import (
	"net/http"
	"testing"

	"github.com/Malwarize/webpalm/v2/core"
)

func TestCrawler_Crawl(t *testing.T) {
	crawler := core.Crawler{
		RootURL:      "file://arabian_nights.html",
		Level:        2,
		Client:       &http.Client{},
		LiveMode:     true,
		ExportFile:   "test.xml",
		IncludedUrls: []string{"youtube.com"},
		RegexMap: map[string]string{
			"email":    "[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+",
			"comments": "<!--.*?-->",
		},
	}
	crawler.Crawl()
}
