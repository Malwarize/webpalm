package core

import (
	"github.com/XORbit01/webpalm/webtree"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	GeneralRegex = regexp.MustCompile(`((https?://)?([\da-z.-]+\.[a-z]{2,6}|[\da-z.-]+\.[a-z]{2,6}\.[a-z]{2,6})(/[\w.-]*)*/?)`)
)

type Crawler struct {
	RootURL string
	Level   int
	Client  *http.Client
}

func (c *Crawler) Crawl() {
	root := webtree.Node{}
	root.Page.SetUrl(c.RootURL)
	c.CrawlNode(&root, c.Level)
	root.Display()
}

func (c *Crawler) Fetch(page *webtree.Page) {
	if strings.HasPrefix(page.GetUrl(), "file://") {
		// only for testing purposes
		data, err := os.ReadFile(page.GetUrl()[7:])
		if err != nil {
			return
		}
		page.SetData(string(data))
		return
	}
	req, err := http.NewRequest("GET", page.GetUrl(), nil)
	if err != nil {
		return
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	page.SetData(string(body))
	page.SetStatusCode(resp.StatusCode)
}

func (c *Crawler) ExtractLinks(page *webtree.Page) (links []string) {
	matches := GeneralRegex.FindAllString(page.GetData(), -1)
	for _, link := range matches {
		links = append(links, link)
	}
	return
}

func (c *Crawler) CrawlNode(w *webtree.Node, level int) {
	var f func(w *webtree.Node, level int, maxDepth int)
	f = func(w *webtree.Node, level int, maxDepth int) {
		if level >= maxDepth {
			return
		}
		c.Fetch(&w.Page)
		links := c.ExtractLinks(&w.Page)
		for _, link := range links {
			child := w.AddChild(webtree.Page{})
			child.Page.SetUrl(link)
			f(child, level+1, maxDepth)
		}
	}
	f(w, 1, level)
}
