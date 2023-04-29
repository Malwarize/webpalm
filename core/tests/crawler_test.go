package tests

import (
	"github.com/XORbit01/webpalm/core"
	"net/http"
	"testing"
)

func TestCrawler_Crawl(t *testing.T) {
	crawler := core.Crawler{
		RootURL:    "file://arabian_nights.html",
		Level:      1,
		Client:     &http.Client{},
		LiveMode:   false,
		ExportFile: "test.xml",
		RegexMap: map[string]string{
			"email":    "[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+",
			"comments": "<!--.*?-->",
		},
	}
	crawler.Crawl()
}
