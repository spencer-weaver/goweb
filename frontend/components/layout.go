package components

import (
	"github.com/spencer-weaver/goweb/internal/template/html"
)

func HeroSection(file string) *html.Component {
	section := html.NewSection()
	section.AddAttribute("class", "hero")

	section.AddChild(Markdown(file))

	return section
}

func AboutSection(file string) *html.Component {
	section := html.NewSection()
	section.AddAttribute("class", "about")
	section.AddChild(Markdown(file))
	return section
}
