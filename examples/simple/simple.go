//go:build go1.18
// +build go1.18

package main

import (
	"net/http"

	. "github.com/melias122/html"
)

func main() {
	_ = http.ListenAndServe("localhost:8080", http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	_ = Page(props{
		title: r.URL.Path,
		path:  r.URL.Path,
	}).Render(w)
}

type props struct {
	title string
	path  string
}

// Page is a whole document to output.
func Page(p props) Node {
	return HTML5(HTML5Props{
		Title:    p.title,
		Language: "en",
		Head: []Node{
			StyleEl(Type("text/css"),
				Raw("html { font-family: sans-serif; }"),
				Raw("ul { list-style-type: none; margin: 0; padding: 0; overflow: hidden; }"),
				Raw("ul li { display: block; padding: 8px; float: left; }"),
				Raw(".is-active { font-weight: bold; }"),
			),
		},
		Body: []Node{
			Navbar(p.path, []PageLink{
				{Path: "/foo", Name: "Foo"},
				{Path: "/bar", Name: "Bar"},
			}),
			H1(Text(p.title)),
			P(Textf("Welcome to the page at %v.", p.path)),
		},
	})
}

type PageLink struct {
	Path string
	Name string
}

func Navbar(currentPath string, links []PageLink) Node {
	return Div(
		Ul(
			NavbarLink("/", "Home", currentPath),

			Group(Map(links, func(pl PageLink) Node {
				return NavbarLink(pl.Path, pl.Name, currentPath)
			})),
		),

		Hr(),
	)
}

func NavbarLink(href, name, currentPath string) Node {
	return Li(A(Href(href), Classes{"is-active": currentPath == href}, Text(name)))
}
