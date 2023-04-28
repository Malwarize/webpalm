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
func (page *Page) PrintPageLive(prefix *string, last bool) {
	fmt.Print(*prefix)
	if last {
		fmt.Print("└── ")
		*prefix += "    "
	} else {
		fmt.Print("├── ")
		*prefix += "│   "
	}
	fmt.Printf("%s (%d)\n", page.GetUrl(), page.GetStatusCode())
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
	fmt.Printf("%s", prefix)
	if isLast {
		fmt.Printf("└── ")
		prefix += "    "
	} else {
		fmt.Printf("├── ")
		prefix += "│   "
	}
	fmt.Printf("%s (%d)\n", node.Page.GetUrl(), node.Page.GetStatusCode())

	for i, child := range node.Children {
		isLast := i == len(node.Children)-1
		child.printTree(prefix, isLast)
	}
}
func (node *Node) Display() {
	node.printTree("", true)
}
