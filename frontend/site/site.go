package site

import (
	"github.com/spencer-weaver/goweb/frontend/components"
	"github.com/spencer-weaver/goweb/internal/template/html"
)

var Pages = map[string]*html.Component{
	PageNames[0]: PageComponents[0],
	PageNames[1]: PageComponents[1],
}

var PageComponents = []*html.Component{
	Home(),
	Models(),
}

var PageNames = []string{
	"Home",
	"Models",
}

func Home() *html.Component {

	page := components.Page()

	document := page.GetFirstElement("html")
	document.AddChild(components.Head("Home"))

	head := document.GetFirstElement("head")
	head.AddChild(components.StyleSheet("./styles/reset.css"))

	nav := components.Nav(PageNames, components.ToTitle)

	body := document.AddChild(components.Body(components.Header(PageNames, nav), html.NewMain(), html.NewFooter()))

	main := body.GetFirstElement("main")

	main.AddChild(components.HeroSection("frontend/markdown/home/hero.md"))
	main.AddChild(components.AboutSection("frontend/markdown/home/about.md"))

	main.AddButton()

	return page
}

func Models() *html.Component {

	page := components.Page()

	document := page.GetFirstElement("html")
	document.AddChild(components.Head("Models"))

	head := document.GetFirstElement("head")
	head.AddChild(components.StyleSheet("./styles/reset.css"))

	nav := components.Nav(PageNames, components.ToTitle)

	body := document.AddChild(components.Body(components.Header(PageNames, nav), html.NewMain(), html.NewFooter()))

	main := body.GetFirstElement("main")
	main.AddChild(components.Button("button"))

	return page
}
