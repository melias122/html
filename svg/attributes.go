package svg

import (
	"github.com/melias122/html"
)

func ClipRule(v string) html.Node {
	return html.Attr("clip-rule", v)
}

func D(v string) html.Node {
	return html.Attr("d", v)
}

func Fill(v string) html.Node {
	return html.Attr("fill", v)
}

func FillRule(v string) html.Node {
	return html.Attr("fill-rule", v)
}

func Stroke(v string) html.Node {
	return html.Attr("stroke", v)
}

func ViewBox(v string) html.Node {
	return html.Attr("viewBox", v)
}
