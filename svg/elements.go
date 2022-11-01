// Package svg provides common SVG elements and attributes.
// See https://developer.mozilla.org/en-US/docs/Web/SVG/Element for an overview.
package svg

import (
	"github.com/melias122/html"
)

func Path(children ...html.Node) html.Node {
	return html.El("path", children...)
}

func SVG(children ...html.Node) html.Node {
	return html.El("svg", html.Attr("xmlns", "http://www.w3.org/2000/svg"), html.Group(children))
}
