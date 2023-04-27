package core

import (
	"github.com/XORbit01/webpalm/webtree"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
)

var (
	GeneralRegex = `((?:https?)://[\w\-]+(?:\.[\w\-]+)+[\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])`
)

type Crawler struct {
	RootURL string
	Level   int
	Client  *http.Client
}

func NewCrawler(url string, level int) *Crawler {
	return &Crawler{
		RootURL: url,
		Level:   level,
		Client:  &http.Client{},
	}
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
		page.SetStatusCode(200)
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
	regex := regexp.MustCompile(GeneralRegex)
	matches := regex.FindAllString(page.GetData(), -1)
	for _, link := range matches {
		links = append(links, link)
	}
	return
}

func (c *Crawler) CrawlNode(w *webtree.Node, level int) {
	if level < 0 {
		return
	}
	//leaf nodes
	if level == 0 {
		c.Fetch(&w.Page)
		return
	}
	c.Fetch(&w.Page)
	links := c.ExtractLinks(&w.Page)
	if w.Page.GetStatusCode() == 0 {
		return
	}
	// add children
	wg := sync.WaitGroup{}
	for _, link := range links {
		wg.Add(1)
		go func(link string) {
			child := w.AddChild(webtree.Page{})
			child.Page.SetUrl(link)
			defer wg.Done()
		}(link)
	}
	wg.Wait()

	// crawl children
	for _, child := range w.Children {
		wg.Add(1)
		go func(child *webtree.Node, level int) {
			defer wg.Done()
			c.CrawlNode(child, level)
		}(child, level-1)
	}
	wg.Wait()
}
