// Package components provides high-level components and helpers that are composed of low-level elements and attributes.
package html

// HTML5Props for HTML5.
// Title is set no matter what, Description and Language elements only if the strings are non-empty.
type HTML5Props struct {
	Title       string
	Description string
	Language    string
	Head        []Node
	Body        []Node
}

// HTML5 document template.
func HTML5(p HTML5Props) Node {
	return Doctype(
		HTML(If(p.Language != "", Lang(p.Language)),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
				TitleEl(Text(p.Title)),
				If(p.Description != "", Meta(Name("description"), Content(p.Description))),
				Group(p.Head),
			),
			Body(Group(p.Body)),
		),
	)
}
