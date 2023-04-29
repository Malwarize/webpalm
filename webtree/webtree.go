package webtree

import (
	"fmt"
)

type Node struct {
	Page     Page
	Parent   *Node
	Children []*Node
}

func (node *Node) AddChild(page *Page) *Node {
	child := &Node{Page: *page, Parent: node}
	node.Children = append(node.Children, child)
	return child
}

func (node *Node) SprintJSON() ([]byte, error) {
	return node.ToJSONPage().SprintJSON()
}
func (node *Node) SprintTXT() (string, error) {
	var out string = ""
	var f func(node *Node, prefix string, isLast bool)
	f = func(node *Node, prefix string, isLast bool) {
		out += node.Page.SprintPageLine(&prefix, isLast)
		for i, child := range node.Children {
			isLast := i == len(node.Children)-1
			f(child, prefix, isLast)
		}
	}
	f(node, "", true)
	return out, nil
}
func (node *Node) SprintTXTColored() (string, error) {
	var out string = ""
	var f func(node *Node, prefix string, isLast bool)
	f = func(node *Node, prefix string, isLast bool) {
		out += node.Page.SprintPageLineColored(&prefix, isLast)
		for i, child := range node.Children {
			isLast := i == len(node.Children)-1
			f(child, prefix, isLast)
		}
	}
	f(node, "", true)
	return out, nil
}
func (node *Node) SprintXML() ([]byte, error) {
	return node.ToXMLPage().SprintXML()
}

func (node *Node) ToJSONPage() *JsonPage {
	exportNode := NewJsonPage()
	exportNode.URL = node.Page.GetUrl()
	exportNode.StatusCode = node.Page.GetStatusCode()
	for name, results := range node.Page.GetResults() {
		for _, result := range results {
			exportNode.Results[name] = append(exportNode.Results[name], result)
		}
	}
	exportNode.Children = make([]*JsonPage, 0)
	for _, child := range node.Children {
		exportNode.Children = append(exportNode.Children, child.ToJSONPage())
	}
	return exportNode
}

func (node *Node) ToXMLPage() *XmlPage {
	exportNode := NewXmlPage()
	exportNode.URL = node.Page.GetUrl()
	exportNode.StatusCode = node.Page.GetStatusCode()
	for name, results := range node.Page.GetResults() {
		for _, result := range results {
			exportNode.Results = append(exportNode.Results, &XmlPageResult{Pattern: name, Result: []string{result}})
		}
	}
	exportNode.Children = make([]*XmlPage, 0)
	for _, child := range node.Children {
		exportNode.Children = append(exportNode.Children, child.ToXMLPage())
	}
	return exportNode
}

func (node *Node) Display() {
	out, err := node.SprintTXTColored()
	if err != nil {
		fmt.Println(out)
		return
	}
	fmt.Print(out)
}
