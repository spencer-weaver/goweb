package components

import "github.com/spencer-weaver/goweb/internal/template/html"

func Markdown(file string) *html.Component {
	markdown := html.NewComponent()
	markdown.Render.FilePath = file
	markdown.Render.Type = "markdown"
	return markdown
}
