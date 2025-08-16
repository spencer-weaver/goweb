package components

import "github.com/spencer-weaver/goweb/internal/template/html"

func Button(buttonType string) *html.Component {
	button := html.NewButton()

	button.AddAttribute("type", buttonType)

	return button
}
