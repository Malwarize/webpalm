package webtree

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/fatih/color"
	"net/url"
	"path"
)

type Page struct {
	url        string
	statusCode int
	data       string
	results    map[string][]string
}

type JsonPage struct {
	URL        string              `json:"url"`
	StatusCode int                 `json:"status_code"`
	Results    map[string][]string `json:"results"`
	Children   []*JsonPage         `json:"children,omitempty"`
}

type XmlPage struct {
	URL        string           `xml:"url"`
	StatusCode int              `xml:"status_code"`
	Results    []*XmlPageResult `xml:"results"`
	Children   []*XmlPage       `xml:"children,omitempty"`
}
type XmlPageResult struct {
	Pattern string   `xml:"pattern"`
	Result  []string `xml:"result"`
}

func NewPage() *Page {
	return &Page{
		results: make(map[string][]string),
	}
}
func NewJsonPage() *JsonPage {
	return &JsonPage{
		Results: make(map[string][]string),
	}
}
func NewXmlPage() *XmlPage {
	return &XmlPage{
		Results: make([]*XmlPageResult, 0),
	}
}

func (p XmlPage) SprintXML() ([]byte, error) {
	data, err := xml.MarshalIndent(p, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (p JsonPage) SprintJSON() ([]byte, error) {
	data, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (page *Page) GetUrl() string {
	return page.url
}
func (page *Page) SetUrl(url string) {
	page.url = url
}

func (page *Page) GetStatusCode() int {
	return page.statusCode
}
func (page *Page) SetStatusCode(code int) {
	page.statusCode = code
}

func (page *Page) GetData() string {
	return page.data
}
func (page *Page) SetData(s string) {
	page.data = s
}

func (page *Page) SprintPageLineColored(prefix *string, last bool) string {
	out := fmt.Sprint(*prefix)
	if last {
		out += color.BlueString("└── ")
		*prefix += "    "
	} else {
		out += color.BlueString("├── ")
		*prefix += color.BlueString("│   ")
	}

	out += color.RedString("[")
	out += fmt.Sprintf("%s", page.GetUrl()) // [google.com]
	out += color.RedString("]")

	out += color.CyanString("(")
	out += color.CyanString("%d", page.GetStatusCode()) // (200)
	out += color.CyanString(")\n")
	/*
		|- [google.com] (200)
		    ┌
		    │ google
		    │ google.com
		    └
	*/
	for match, results := range page.GetResults() {
		out += fmt.Sprintf(*prefix)
		out += fmt.Sprintf("  %s\n", match) // emails :

		out += fmt.Sprintf(*prefix)
		out += color.RedString("    ┌\n") //	┌
		for _, result := range results {
			out += fmt.Sprintf(*prefix)
			out += color.RedString("    │")
			out += color.GreenString("%s\n", result) // | example@gmail.com
		}
		out += fmt.Sprintf(*prefix)
		out += color.RedString("    └\n") // └

	}
	return out
}
func (page *Page) SprintPageLine(prefix *string, last bool) string {
	out := fmt.Sprint(*prefix)
	if last {
		out += "└── "
		*prefix += "    "
	} else {
		out += "├── "
		*prefix += "│   "
	}

	out += "["
	out += fmt.Sprintf("%s", page.GetUrl()) // [google.com]
	out += "]"

	out += "("
	out += fmt.Sprintf("%d", page.GetStatusCode()) // (200)
	out += ")\n"
	/*
		|- [google.com] (200)
		    ┌
		    │ google
		    │ google.com
		    └
	*/
	for match, results := range page.GetResults() {
		out += fmt.Sprintf(*prefix)
		out += fmt.Sprintf("  %s\n", match) // emails :

		out += fmt.Sprintf(*prefix)
		out += "    ┌\n" //	┌
		for _, result := range results {
			out += fmt.Sprintf(*prefix)
			out += "    │"
			out += fmt.Sprintf("%s\n", result) // |
		}
		out += fmt.Sprintf(*prefix)
		out += "    └\n" // └
	}
	return out
}

func (page *Page) PrintPageLive(prefix *string, last bool) {
	fmt.Print(page.SprintPageLineColored(prefix, last))
}

func (page *Page) AddMatch(rname string, match string) {
	if page.results == nil {
		page.results = make(map[string][]string)
	}
	if _, ok := page.results[rname]; !ok {
		page.results[rname] = make([]string, 0)
	}
	page.results[rname] = append(page.results[rname], match)
}
func (page *Page) GetResults() map[string][]string {
	return page.results
}
func (page *Page) ConvertToAbsoluteURL(relativePath string) (string, error) {
	base, err := url.Parse(page.GetUrl())
	if err != nil {
		return "", err
	}

	absoluteUrl := base.ResolveReference(&url.URL{Path: relativePath})
	absolutePath := absoluteUrl.Path
	if !path.IsAbs(absolutePath) {
		absolutePath = path.Join(base.Path, absolutePath)
	}
	absoluteUrl.Path = absolutePath
	return absoluteUrl.String(), nil
}
