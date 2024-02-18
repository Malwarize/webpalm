package core_tests

import (
	"net/http"
	"testing"

	"github.com/Malwarize/webpalm/v2/core"
	"github.com/stretchr/testify/assert"
)

func TestCrawler(t *testing.T) {
	var exportFile = "test.xml"
	crawler := core.Crawler{
		RootURL:      "file://arabian_nights.html",
		Level:        2,
		Client:       &http.Client{},
		ExportFile:   exportFile,
		IncludedUrls: []string{"youtube.com"},
		RegexMap: map[string]string{
			"email":    "[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+",
			"comments": "<!--.*?-->",
		},
	}
	crawler.Crawl()

	assert.FileExists(t, exportFile, "File exists.")

}
