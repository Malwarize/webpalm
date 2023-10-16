package core

import (
	"crypto/tls"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"io"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/Malwarize/webpalm/v2/shared"
	"github.com/Malwarize/webpalm/v2/webtree"
)

var (
	GeneralRegex = `((?:https?)://[\w\-]+(?:\.[\w\-]+)+[\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])`
	HrefRegex    = `href=["']([^"']+)["']`
)

var UnreadableExtensions = []string{
	".png",
	".jpg",
	".jpeg",
	".gif",
	".pdf",
	".doc",
	".docx",
	".xls",
	".xlsx",
	".ppt",
	".pptx",
	".zip",
	".rar",
	".tar",
	".gz",
	".exe",
	".mp3",
	".mp4",
	".avi",
	".mov",
	".wmv",
	".flv",
	".wav",
	".mpeg",
	".mpg",
	".m4v",
	".swf",
	".svg",
	".ico",
	".ttf",
	".woff",
	".woff2",
	".eot",
	".otf",
	".psd",
	".ai",
	".eps",
	".indd",
	".raw",
	".webm",
	".m4a",
	".m4p",
	".m4b",
	".m4r",
}

type Crawler struct {
	RootURL        string
	Level          int
	ExportFile     string
	RegexMap       map[string]string
	ExcludedStatus []int
	IncludedUrls   []string
	Client         *http.Client
	UserAgent      string
	Cache          Cache
	Workers        int
	Delay          int
}

func NewCrawler(options *shared.Options) *Crawler {
	crawler := Crawler{
		RootURL:        options.URL,
		Level:          options.Level,
		ExportFile:     options.ExportFile,
		RegexMap:       options.RegexMap,
		ExcludedStatus: options.StatusResponses,
		IncludedUrls:   options.IncludedUrls,
		Workers:        options.Workers,
		Delay:          options.Delay,
		UserAgent:      options.UserAgent,
		Cache: Cache{
			Visited: make(map[string]bool),
		},
	}
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	if options.Proxy != nil {
		transport.Proxy = http.ProxyURL(options.Proxy)
	}
	crawler.Client = &http.Client{
		Transport: transport,
		Timeout:   time.Duration(options.TimeOut) * time.Second,
	}
	return &crawler
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
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		page.SetStatusCode(-1)
		page.SetData("")
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		page.SetStatusCode(resp.StatusCode)
		page.SetData("")
		return
	}
	page.SetData(string(body))
	page.SetStatusCode(resp.StatusCode)
}

func (c *Crawler) ExtractLinks(page *webtree.Page) (links []string) {
	regex := regexp.MustCompile(GeneralRegex)
	generalUrlMatches := regex.FindAllString(page.GetData(), -1)
	links = append(links, generalUrlMatches...)
	hrefRegex := regexp.MustCompile(HrefRegex)
	hrefMatches := hrefRegex.FindAllStringSubmatch(page.GetData(), -1)
	for _, match := range hrefMatches {
		// check if it is a normal url
		if strings.HasPrefix(match[1], "http") ||
			strings.HasPrefix(match[1], "https") ||
			strings.HasPrefix(match[1], "file") {
			links = append(links, match[1])
			continue
		}
		// check if it is a relative url
		if strings.HasPrefix(match[1], "/") || strings.HasPrefix(match[1], "./") || strings.HasPrefix(match[1], "../") || strings.HasSuffix(match[1], "/") {
			u, err := page.ConvertToAbsoluteURL(match[1])
			if err != nil {
				continue
			}
			links = append(links, u)
		}
	}
	return
}

func (c *Crawler) ExportJSON(root webtree.Node, filename string) error {
	data, err := root.SprintJSON()
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0o644)
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
	err = os.WriteFile(filename, []byte(data), 0o644)
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
	err = os.WriteFile(filename, data, 0o644)
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

func (c *Crawler) isSkipableUrl(u string) bool {
	for _, v := range UnreadableExtensions {
		if strings.HasSuffix(u, v) {
			return true
		}
	}
	// get domain name from url
	if strings.Contains(c.RootURL, u) {
		return false
	}
	if len(c.IncludedUrls) == 0 {
		return false
	}
	for _, v := range c.IncludedUrls {
		if strings.Contains(u, v) {
			return false
		}
	}
	return true
}

func (c *Crawler) IsSkipablePage(page webtree.Page) bool {
	isInCode := func(status int, arr []int) bool {
		for _, v := range arr {
			if v == status {
				return true
			}
		}
		return false
	}
	if page.GetStatusCode() == 0 ||
		isInCode(page.GetStatusCode(), c.ExcludedStatus) ||
		c.isSkipableUrl(page.GetUrl()) ||
		c.Cache.IsVisited(page.GetUrl()) {
		return true
	}
	return false
}

func (c *Crawler) AddMatches(page webtree.Page) {
	for rname, regex := range c.RegexMap {
		r := regexp.MustCompile(regex)
		matches := r.FindAllString(page.GetData(), -1)
		for _, match := range matches {
			page.AddMatch(rname, match)
		}
	}
}

func (c *Crawler) ProcessANode(node *webtree.Node) {
	c.Fetch(&node.Page)
	c.AddMatches(node.Page)
	if c.IsSkipablePage(node.Page) {
		return
	}
	c.Cache.AddVisited(node.Page.GetUrl())
	if c.Level < 1 {
		return
	}
	links := c.ExtractLinks(&node.Page)
	for _, link := range links {
		if c.isSkipableUrl(link) {
			continue
		}
		child := node.AddChild(webtree.NewPage())
		child.Page.SetUrl(link)
	}
}

func (c *Crawler) worker(tasks <-chan *webtree.Node, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		c.ProcessANode(task)
	}
}

func (c *Crawler) CrawlNodeBlock(w *webtree.Node, levelChangedChan chan int) {
	c.ProcessANode(w)
	level := c.Level
	for c.Level >= 0 {
		var wg sync.WaitGroup
		tasks := make(chan *webtree.Node, c.Workers)
		for i := 0; i < c.Workers; i++ {
			wg.Add(1)
			go c.worker(tasks, &wg)
		}
		children := w.GetAllChildrenOfLevel(level - c.Level)
		for _, child := range children {
			tasks <- child
		}
		close(tasks)
		wg.Wait()

		//signal to spinner that level has changed
		c.Level--
		levelChangedChan <- level - c.Level
	}
}

func (c *Crawler) CrawlNodeLive(w *webtree.Node) {
	var f func(w *webtree.Node, level int, prefix string, last bool)
	f = func(w *webtree.Node, level int, prefix string, last bool) {
		if level < 0 {
			return
		}
		c.Fetch(&w.Page)

		if c.Delay > 0 {
			time.Sleep(time.Duration(c.Delay) * time.Millisecond)
		}

		// add matches
		c.AddMatches(w.Page)

		if c.IsSkipablePage(w.Page) {
			return
		}
		w.Page.PrintPageLive(&prefix, last)

		// leaf nodes
		if level == 0 {
			return
		}
		// add visited node to cache
		c.Cache.AddVisited(w.Page.GetUrl())

		links := c.ExtractLinks(&w.Page)

		// add children
		for i, link := range links {
			if c.isSkipableUrl(link) {
				continue
			}
			child := w.AddChild(webtree.NewPage())
			child.Page.SetUrl(link)
			f(child, level-1, prefix, i == len(links)-1)
		}
	}
	f(w, c.Level, "", true)
}

func (c *Crawler) SaveResults(root webtree.Node) {
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
		err := c.Export(root, "json", c.ExportFile)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (c *Crawler) Crawl() {
	root := webtree.Node{}
	root.Page.SetUrl(c.RootURL)
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-interruptChan
		fmt.Println("\033[?25h")
		if c.Workers > 0 {
			root.Display()
		}
		if c.ExportFile != "" {
			fmt.Println("Saving results to file...")
			c.SaveResults(root)
		}
		c.Cache.Flush()
		os.Exit(0)
	}()

	// live mode or block mode
	if c.Workers == 0 {
		c.CrawlNodeLive(&root)
	} else {
		color.Yellow("NOTE: This program is running in parallel mode, so you won't be able to see the output directly until the tree is entirely built. You can observe the output in your saved file, which is updated at each level traversal.")
		LevelChangedChan := make(chan int, 1)
		go func() {
			s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
			s.Prefix = color.GreenString("incursion ... ")
			s.Suffix = color.CyanString(" [ crawling level %d ]", 0)
			_ = s.Color("yellow")
			s.Start()
			for {
				l := <-LevelChangedChan
				if c.Level < 0 {
					s.Stop()
					return
				}
				s.Suffix = color.CyanString(" [ crawling level %d ]", l)
				s.Restart()
				// save results
				if c.ExportFile != "" {
					c.SaveResults(root)
				}
			}
		}()

		c.CrawlNodeBlock(&root, LevelChangedChan)
		root.Display()
	}
	fmt.Println("\033[?25h")
	if c.ExportFile != "" {
		c.SaveResults(root)
	}
}
