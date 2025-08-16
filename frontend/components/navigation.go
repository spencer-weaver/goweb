package components

import (
	"strings"

	"github.com/spencer-weaver/goweb/internal/template/html"
)

func Nav(pages []string, textFunc func(s string) string) *html.Component {
	nav := html.NewNav()

	// links
	ul := nav.AddUl()
	for _, page := range pages {
		link := ul.AddLi().AddA()
		link.AddAttribute("href", "./"+strings.ToLower(page)+".html")
		link.InnerHTML = string(textFunc(page))
	}

	return nav
}
