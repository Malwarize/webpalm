package core

import (
	"fmt"
	"github.com/XORbit01/webpalm/webtree"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	GeneralRegex = `((?:https?)://[\w\-]+(?:\.[\w\-]+)+[\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])`
)

type Crawler struct {
	RootURL    string
	Level      int
	OutputMode string
	Client     *http.Client
}

func NewCrawler(url string, level int, outputMode string) *Crawler {
	return &Crawler{
		RootURL:    url,
		Level:      level,
		OutputMode: outputMode,
		Client:     &http.Client{},
	}
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

func (c *Crawler) CrawlNodeBlock(w *webtree.Node) {
	var f func(w *webtree.Node, level int)
	f = func(w *webtree.Node, level int) {
		if level < 0 {
			return
		}
		c.Fetch(&w.Page)
		if level == 0 {
			return
		}
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
				f(child, level-1)
				defer wg.Done()
			}(link)
		}
		wg.Wait()
	}
	f(w, c.Level)
}

func (c *Crawler) CrawlNodeLive(w *webtree.Node) {
	var f func(w *webtree.Node, level int, prefix string, last bool)
	f = func(w *webtree.Node, level int, prefix string, last bool) {
		if level < 0 {
			return
		}
		c.Fetch(&w.Page)

		w.Page.PrintPageLive(&prefix, last)
		//leaf nodes
		if level == 0 {
			return
		}
		links := c.ExtractLinks(&w.Page)
		if w.Page.GetStatusCode() == 0 {
			return
		}
		// add children
		for i, link := range links {
			child := w.AddChild(webtree.Page{})
			child.Page.SetUrl(link)
			f(child, level-1, prefix, i == len(links)-1)
		}
	}
	f(w, c.Level, "", true)
}

func (c *Crawler) Crawl() {
	root := webtree.Node{}
	root.Page.SetUrl(c.RootURL)
	// live mode or block mode
	if c.OutputMode == "live" {
		now := time.Now()
		c.CrawlNodeLive(&root)
		fmt.Println("took : ", time.Since(now))
	} else if c.OutputMode == "block" {
		now := time.Now()
		c.CrawlNodeBlock(&root)
		root.Display()
		fmt.Println("took : ", time.Since(now))
	} else {

	}
}
