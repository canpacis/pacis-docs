package view

import (
	"fmt"

	"github.com/canpacis/pacis/html"
)

func Name(name string) *html.Attribute {
	return html.StyleAttr(fmt.Sprintf("view-transition-name: %s", name))
}
