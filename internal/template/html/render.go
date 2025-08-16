package html

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/yuin/goldmark"
)

func LoadHtmlTemplate(file string) *template.Template {
	buffer, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("error reading file %s\n", file)
	}
	tmpl, err := template.New("html-element").Parse(string(buffer))
	if err != nil {
		fmt.Printf("error parsing template file %s\n", file)
	}
	return tmpl
}

func RenderComponent(component *Component) string {

	var componentString string

	innerHtml := ""
	for _, child := range component.Children {
		innerHtml += RenderComponent(child)
	}

	component.InnerHTML += innerHtml

	switch component.Render.Type {
	case "root":
		return innerHtml
	case "html":
		componentString = RenderHTML(component)
	case "markdown":
		componentString = RenderMD(component) + componentString
	default:
		return ""
	}

	return componentString
}

func RenderHTML(component *Component) string {

	tmpl := LoadHtmlTemplate("internal/template/html/templates/elements.html")

	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, component.Render.Template, component)
	if err != nil {
		fmt.Printf("error executing template for component %s: %s\n", component.Name, err)
	}
	return buf.String()
}

func RenderMD(component *Component) string {

	buf, err := os.ReadFile(component.Render.FilePath)
	if err != nil {
		fmt.Printf("error reading file %s for component\n", component.Render.FilePath)
	}

	var htmlBuf bytes.Buffer
	err = goldmark.Convert(buf, &htmlBuf)
	if err != nil {
		fmt.Printf("error converting markdown to html in file %s\n", component.Render.FilePath)
	}

	return htmlBuf.String()
}
