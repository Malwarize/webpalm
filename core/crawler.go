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
)

var (
	GeneralRegex = `((?:https?)://[\w\-]+(?:\.[\w\-]+)+[\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])`
)

type Crawler struct {
	RootURL        string
	Level          int
	LiveMode       bool
	ExportFile     string
	RegexMap       map[string]string
	ExcludedStatus []int
	Client         *http.Client
}

func NewCrawler(url string, level int, liveMode bool, exportFile string, regexMap map[string]string, statusResponses []int) *Crawler {
	return &Crawler{
		RootURL:        url,
		Level:          level,
		LiveMode:       liveMode,
		ExportFile:     exportFile,
		RegexMap:       regexMap,
		ExcludedStatus: statusResponses,
		Client:         &http.Client{},
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
		// add matches
		func() {
			for rname, regex := range c.RegexMap {
				r := regexp.MustCompile(regex)
				matches := r.FindAllString(w.Page.GetData(), -1)
				for _, match := range matches {
					w.Page.AddMatch(rname, match)
				}
			}
		}()

		isIn := func(status int, arr []int) bool {
			for _, v := range arr {
				if v == status {
					return true
				}
			}
			return false
		}
		if isIn(w.Page.GetStatusCode(), c.ExcludedStatus) || w.Page.GetStatusCode() == 0 {
			return
		}

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
				child := w.AddChild(webtree.NewPage())
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
		func() {
			for rname, regex := range c.RegexMap {
				r := regexp.MustCompile(regex)
				matches := r.FindAllString(w.Page.GetData(), -1)
				for _, match := range matches {
					w.Page.AddMatch(rname, match)
				}
			}
		}()
		isIn := func(status int, arr []int) bool {
			for _, v := range arr {
				if v == status {
					return true
				}
			}
			return false
		}
		if isIn(w.Page.GetStatusCode(), c.ExcludedStatus) || w.Page.GetStatusCode() == 0 {
			return
		}
		w.Page.PrintPageLive(&prefix, last)
		//leaf nodes
		if level == 0 {
			return
		}
		links := c.ExtractLinks(&w.Page)

		// add children
		for i, link := range links {
			child := w.AddChild(webtree.NewPage())
			child.Page.SetUrl(link)
			f(child, level-1, prefix, i == len(links)-1)
		}
	}
	f(w, c.Level, "", true)
}

func (c *Crawler) ExportJSON(root webtree.Node, filename string) error {
	data, err := root.SprintJSON()
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (c *Crawler) ExportTXT(root webtree.Node, filename string) error {
	data, err := root.SprintTXT()
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (c *Crawler) ExportXML(tree webtree.Node, filename string) error {
	data, err := tree.SprintXML()
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (c *Crawler) Export(tree webtree.Node, format string, filename string) error {
	if format == "json" {
		err := c.ExportJSON(tree, filename)
		if err != nil {
			return err
		}
	}
	if format == "txt" {
		err := c.ExportTXT(tree, filename)
		if err != nil {
			return err
		}
	}
	if format == "xml" {
		err := c.ExportXML(tree, filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Crawler) Crawl() {
	root := webtree.Node{}
	root.Page.SetUrl(c.RootURL)
	// live mode or block mode
	if c.LiveMode {
		c.CrawlNodeLive(&root)
	} else {
		c.CrawlNodeBlock(&root)
		root.Display()
	}
	if c.ExportFile != "" {
		if strings.HasSuffix(c.ExportFile, ".txt") {
			err := c.Export(root, "txt", c.ExportFile)
			if err != nil {
				fmt.Println(err)
			}
		} else if strings.HasSuffix(c.ExportFile, ".xml") {
			err := c.Export(root, "xml", c.ExportFile)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			//default to json
			err := c.Export(root, "json", c.ExportFile)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
