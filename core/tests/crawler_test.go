package tests

import (
	"github.com/XORbit01/webpalm/core"
	"net/http"
	"testing"
)

func TestCrawler_Crawl(t *testing.T) {
	crawler := core.Crawler{
		RootURL:    "file://arabian_nights.html",
		Level:      2,
		Client:     &http.Client{},
		OutputMode: "live",
	}
	crawler.Crawl()
}
