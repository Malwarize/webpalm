package webtree

import (
	"reflect"
	"testing"
)

func TestNode_AddChild(t *testing.T) {
	type fields struct {
		Page     Page
		Parent   *Node
		Children []*Node
	}
	type args struct {
		page *Page
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				Page:     tt.fields.Page,
				Parent:   tt.fields.Parent,
				Children: tt.fields.Children,
			}
			if got := node.AddChild(tt.args.page); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddChild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Display(t *testing.T) {
	type fields struct {
		Page     Page
		Parent   *Node
		Children []*Node
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				Page:     tt.fields.Page,
				Parent:   tt.fields.Parent,
				Children: tt.fields.Children,
			}
			node.Display()
		})
	}
}

func TestNode_SprintJSON(t *testing.T) {
	type fields struct {
		Page     Page
		Parent   *Node
		Children []*Node
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				Page:     tt.fields.Page,
				Parent:   tt.fields.Parent,
				Children: tt.fields.Children,
			}
			got, err := node.SprintJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("SprintJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SprintJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_SprintTXT(t *testing.T) {
	type fields struct {
		Page     Page
		Parent   *Node
		Children []*Node
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				Page:     tt.fields.Page,
				Parent:   tt.fields.Parent,
				Children: tt.fields.Children,
			}
			got, err := node.SprintTXT()
			if (err != nil) != tt.wantErr {
				t.Errorf("SprintTXT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SprintTXT() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_SprintTXTColored(t *testing.T) {
	type fields struct {
		Page     Page
		Parent   *Node
		Children []*Node
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				Page:     tt.fields.Page,
				Parent:   tt.fields.Parent,
				Children: tt.fields.Children,
			}
			got, err := node.SprintTXTColored()
			if (err != nil) != tt.wantErr {
				t.Errorf("SprintTXTColored() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SprintTXTColored() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_SprintXML(t *testing.T) {
	type fields struct {
		Page     Page
		Parent   *Node
		Children []*Node
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				Page:     tt.fields.Page,
				Parent:   tt.fields.Parent,
				Children: tt.fields.Children,
			}
			got, err := node.SprintXML()
			if (err != nil) != tt.wantErr {
				t.Errorf("SprintXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SprintXML() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_ToJSONPage(t *testing.T) {
	type fields struct {
		Page     Page
		Parent   *Node
		Children []*Node
	}
	tests := []struct {
		name   string
		fields fields
		want   *JsonPage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				Page:     tt.fields.Page,
				Parent:   tt.fields.Parent,
				Children: tt.fields.Children,
			}
			if got := node.ToJSONPage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToJSONPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_ToXMLPage(t *testing.T) {
	type fields struct {
		Page     Page
		Parent   *Node
		Children []*Node
	}
	tests := []struct {
		name   string
		fields fields
		want   *XmlPage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				Page:     tt.fields.Page,
				Parent:   tt.fields.Parent,
				Children: tt.fields.Children,
			}
			if got := node.ToXMLPage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToXMLPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
