package html

import (
	"strings"
)

type Attribute struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Node struct {
	Name       string       `json:"name"`
	Attributes []*Attribute `json:"attributes"`
	InnerHTML  string       `json:"innerHTML"`
	Children   []*Component `json:"children"`
}

func IsSelfClosing(s string) bool {
	// hard coded list of self closing elements, as there are not many
	// could do better
	var voidElements = map[string]bool{
		"area":   true,
		"base":   true,
		"br":     true,
		"col":    true,
		"embed":  true,
		"hr":     true,
		"img":    true,
		"input":  true,
		"link":   true,
		"meta":   true,
		"source": true,
		"track":  true,
		"wbr":    true,
		"path":   true,
	}
	return voidElements[strings.ToLower(s)]
}

func NewNode(name string) *Node {
	return &Node{
		Name: name,
	}
}

func (component *Component) AddAttribute(label, value string) *Component {
	component.Attributes = append(component.Attributes, &Attribute{
		Label: label,
		Value: value,
	})
	return component
}

func (component *Component) AddClass(class ...string) *Component {
	for _, attr := range component.Attributes {
		if attr.Label == "class" {
			attr.Value += " " + strings.Join(class, " ")
			return component
		}
	}
	component.AddAttribute("class", strings.Join(class, " "))
	return component
}

func (component *Component) AddInlineStyle(class ...string) *Component {
	for _, attr := range component.Attributes {
		if attr.Label == "style" {
			attr.Value += ";" + strings.Join(class, ";")
			return component
		}
	}
	component.AddAttribute("style", strings.Join(class, ";"))
	return component
}

type RenderSpec struct {
	Type        string
	FilePath    string
	Template    string `json:"template"`
	SelfClosing bool   `json:"selfClosing"`
	RenderFunc  func(filePath string) (string, error)
	Status      string `json:"status"`
}

type Component struct {
	Node
	Render RenderSpec
}

func NewComponent() *Component {
	return &Component{}
}

func (component *Component) AddChild(child *Component) *Component {
	component.Children = append(component.Children, child)
	return child
}

func (component *Component) GetFirstElement(name string) *Component {

	queue := []*Component{component}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Name == name {
			return current
		}
		queue = append(queue, current.Children...)
	}
	return nil
}
