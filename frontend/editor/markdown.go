package editor

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spencer-weaver/goweb/frontend/components"
	"github.com/spencer-weaver/goweb/internal/template/html"
)

func fileToString(file string) string {
	buf, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("error reading file %s\n", file)
	}
	return string(buf)
}

func MarkdownEditor(file string) *html.Component {
	page := components.Page()

	doc := page.GetFirstElement("html")

	doc.AddChild(components.Head(filepath.Base(file)))

	header := components.Header([]string{}, components.Nav([]string{}, components.ToTitle))
	header.InnerHTML = filepath.Base(file)

	main := html.NewMain()

	textarea := main.AddTextarea()
	textarea.AddAttribute("id", "editor")
	content := fileToString(file)
	textarea.AddAttribute("placeholder", "Write markdown...")
	if content != "" {
		textarea.InnerHTML = content
	}

	preview := main.AddDiv()
	preview.AddAttribute("id", "preview")

	footer := html.NewFooter()

	cdn := doc.AddScript()
	cdn.AddAttribute("src", "https://cdn.jsdelivr.net/npm/markdown-it/dist/markdown-it.min.js")
	cdn.AddAttribute("defer", "")

	script := doc.AddScript()
	script.InnerHTML = "window.addEventListener('DOMContentLoaded', () => { const md = window.markdownit(); const input = document.getElementById('editor'); const preview = document.getElementById('preview'); input.addEventListener('input', () => { preview.innerHTML = md.render(input.value); }); preview.innerHTML = md.render(input.value); });"
	script.AddAttribute("defer", "")

	doc.AddChild(components.Style("editors/markdown.css"))

	doc.AddChild(components.Body(header, main, footer))

	return page
}
