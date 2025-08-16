package components

import (
	"fmt"
	"os"

	"github.com/spencer-weaver/goweb/internal/template/html"
)

func Style(file string) *html.Component {
	style := html.NewStyle()

	buf, err := os.ReadFile("frontend/styles/" + file)
	if err != nil {
		fmt.Printf("error reading file %s", "frontend/styles/"+file)
	}
	style.InnerHTML = string(buf)

	return style
}

func Page() *html.Component {
	page := html.NewComponent()

	page.Render.Type = "root"

	page.AddDoctype()
	page.AddHtml().AddAttribute("lang", "en")

	return page
}

func Doctype() *html.Component {
	doctype := html.NewDoctype()

	return doctype
}

func Head(name string) *html.Component {
	head := html.NewHead()

	head.AddMeta().AddAttribute("charset", "UTF-8")
	head.AddMeta().AddAttribute("name", "viewport").AddAttribute("content", "width=device-width, initial-scale=1.0")
	head.AddTitle().InnerHTML += name

	return head
}

func Header(pages []string, nav *html.Component) *html.Component {
	header := html.NewHeader()

	header.AddChild(nav)

	return header
}

func Body(header, main, footer *html.Component) *html.Component {
	body := html.NewBody()

	body.AddChild(header)
	body.AddChild(main)
	body.AddChild(footer)

	return body
}

func StyleSheet(path string) *html.Component {
	link := html.NewLink()
	link.AddAttribute("rel", "stylesheet")
	link.AddAttribute("href", path)
	return link
}
