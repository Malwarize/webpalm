package webtree

import "fmt"

type Page struct {
	url        string
	statusCode int
	data       string
}

func (page *Page) GetUrl() string {
	return page.url
}
func (page *Page) GetStatusCode() int {
	return page.statusCode
}
func (page *Page) SetStatusCode(code int) {
	page.statusCode = code
}
func (page *Page) SetUrl(url string) {
	page.url = url
}
func (page *Page) SetData(s string) {
	page.data = s
}

func (page *Page) GetData() string {
	return page.data
}

func (page *Page) Display() {
	println(page.GetUrl())
}

type Node struct {
	Page     Page
	Parent   *Node
	Children []*Node
}

func (node *Node) AddChild(page Page) *Node {
	child := &Node{Page: page, Parent: node}
	node.Children = append(node.Children, child)
	return child
}

func (node *Node) AddChildren(pages []Page) {
	for _, page := range pages {
		node.AddChild(page)
	}
}
func (node *Node) printTree(prefix string, isLast bool) {
	var marker string
	if isLast {
		marker = "└─ "
	} else {
		marker = "├─ "
	}

	if node.Page.GetStatusCode() != 0 {
		fmt.Printf("%s%s%s [%d]\n", prefix, marker, node.Page.url, node.Page.GetStatusCode())
	} else {
		fmt.Printf("%s%s%s\n", prefix, marker, node.Page.url)
	}

	for i, child := range node.Children {
		isLastChild := i == len(node.Children)-1
		var subPrefix string
		if isLastChild {
			subPrefix = prefix + "   "
		} else {
			subPrefix = prefix + "│  "
		}
		child.printTree(subPrefix, isLastChild)
	}
}
func (node *Node) Display() {
	node.printTree("", true)
}
