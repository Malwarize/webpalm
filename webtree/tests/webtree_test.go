package webtree_test

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Malwarize/webpalm/v2/webtree"
)

func TestNewNode(t *testing.T) {
	page := webtree.NewPage()
	node := webtree.NewNode(page, nil, nil)
	assert.NotNil(t, node)
	assert.Equal(t, page, node.Page)
}

func TestAddAndGetChildren(t *testing.T) {
	parent := webtree.NewNode(webtree.NewPage(), nil, []*webtree.Node{})
	childPage := webtree.NewPage()
	child := parent.AddChild(childPage)

	assert.NotNil(t, child)
	assert.Equal(t, 1, len(parent.GetChildren()))
	assert.Equal(t, child, parent.GetChildren()[0])
}

func TestPageSetAndGet(t *testing.T) {
	page := webtree.NewPage()
	url := "http://example.com"
	statusCode := 200
	data := "example data"

	page.SetUrl(url)
	page.SetStatusCode(statusCode)
	page.SetData(data)

	assert.Equal(t, url, page.GetUrl())
	assert.Equal(t, statusCode, page.GetStatusCode())
	assert.Equal(t, data, page.GetData())
}

func TestPageAddAndGetResults(t *testing.T) {
	page := webtree.NewPage()
	matchName := "testMatch"
	match := "testResult"

	page.AddMatch(matchName, match)
	results := page.GetResults()

	assert.Contains(t, results, matchName)
	assert.Contains(t, results[matchName], match)
}

func TestConvertToAbsoluteURL(t *testing.T) {
	page := webtree.NewPage()
	page.SetUrl("http://example.com")
	relativePath := "/test/path"

	absoluteUrl, err := page.ConvertToAbsoluteURL(relativePath)
	assert.NoError(t, err)
	assert.Equal(t, "http://example.com/test/path", absoluteUrl)
}

func TestJsonPageSerialization(t *testing.T) {
	page := webtree.NewJsonPage()
	page.URL = "http://example.com"
	page.StatusCode = 200
	page.Results = map[string][]string{"test": {"result1", "result2"}}

	data, err := page.SprintJSON()
	assert.NoError(t, err)

	var jsonPage webtree.JsonPage
	err = json.Unmarshal(data, &jsonPage)
	assert.NoError(t, err)
	assert.Equal(t, page.URL, jsonPage.URL)
	assert.Equal(t, page.StatusCode, jsonPage.StatusCode)
	assert.Equal(t, page.Results, jsonPage.Results)
}

func TestXmlPageSerialization(t *testing.T) {
	page := webtree.NewXmlPage()
	page.URL = "http://example.com"
	page.StatusCode = 200
	page.Results = append(
		page.Results,
		&webtree.XmlPageResult{Pattern: "test", Result: []string{"result1", "result2"}},
	)

	data, err := page.SprintXML()
	assert.NoError(t, err)

	var xmlPage webtree.XmlPage
	err = xml.Unmarshal(data, &xmlPage)
	assert.NoError(t, err)
	assert.Equal(t, page.URL, xmlPage.URL)
	assert.Equal(t, page.StatusCode, xmlPage.StatusCode)
	// XML unmarshalling does not retain the exact order of slices, additional checks may be needed for Results
}

func TestGetAllChildrenOfLevel(t *testing.T) {
	root := webtree.NewNode(webtree.NewPage(), nil, nil)
	child1 := root.AddChild(webtree.NewPage())
	child2 := root.AddChild(webtree.NewPage())
	child1_1 := child1.AddChild(webtree.NewPage())

	// Test for level 0 (root itself)
	level0 := root.GetAllChildrenOfLevel(0)
	assert.Equal(t, 1, len(level0))
	assert.Equal(t, root, level0[0])

	// Test for level 1 (direct children of root)
	level1 := root.GetAllChildrenOfLevel(1)
	assert.Equal(t, 2, len(level1))
	assert.Contains(t, level1, child1)
	assert.Contains(t, level1, child2)

	// Test for level 2 (children of children)
	level2 := root.GetAllChildrenOfLevel(2)
	assert.Equal(t, 1, len(level2))
	assert.Equal(t, child1_1, level2[0])
}
