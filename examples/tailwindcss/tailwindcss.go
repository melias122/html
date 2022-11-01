//go:build go1.18
// +build go1.18

package main

import (
	"net/http"
	"time"

	. "github.com/melias122/html"
)

func main() {
	http.Handle("/", createHandler(indexPage()))
	http.Handle("/contact", createHandler(contactPage()))
	http.Handle("/about", createHandler(aboutPage()))

	_ = http.ListenAndServe("localhost:8080", nil)
}

func createHandler(title string, body Node) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Rendering a Node is as simple as calling Render and passing an io.Writer
		_ = Page(title, r.URL.Path, body).Render(w)
	}
}

func indexPage() (string, Node) {
	return "Welcome!", Div(
		H1(Text("Welcome to this example page")),
		P(Text("I hope it will make you happy. 😄 It's using TailwindCSS for styling.")),
	)
}

func contactPage() (string, Node) {
	return "Contact", Div(
		H1(Text("Contact us")),
		P(Text("Just do it.")),
	)
}

func aboutPage() (string, Node) {
	return "About", Div(
		H1(Text("About this site")),
		P(Text("This is a site showing off gomponents.")),
	)
}

func Page(title, path string, body Node) Node {
	// HTML5 boilerplate document
	return HTML5(HTML5Props{
		Title:    title,
		Language: "en",
		Head: []Node{
			Link(Rel("stylesheet"), Href("https://unpkg.com/tailwindcss@2.1.2/dist/base.min.css")),
			Link(Rel("stylesheet"), Href("https://unpkg.com/tailwindcss@2.1.2/dist/components.min.css")),
			Link(Rel("stylesheet"), Href("https://unpkg.com/@tailwindcss/typography@0.4.0/dist/typography.min.css")),
			Link(Rel("stylesheet"), Href("https://unpkg.com/tailwindcss@2.1.2/dist/utilities.min.css")),
		},
		Body: []Node{
			Navbar(path, []PageLink{
				{Path: "/contact", Name: "Contact"},
				{Path: "/about", Name: "About"},
			}),
			Container(
				Prose(body),
				PageFooter(),
			),
		},
	})
}

type PageLink struct {
	Path string
	Name string
}

func Navbar(currentPath string, links []PageLink) Node {
	return Nav(Class("bg-gray-700 mb-4"),
		Container(
			Div(Class("flex items-center space-x-4 h-16"),
				NavbarLink("/", "Home", currentPath == "/"),

				// We can Map custom slices to Nodes
				Group(Map(links, func(pl PageLink) Node {
					return NavbarLink(pl.Path, pl.Name, currentPath == pl.Path)
				})),
			),
		),
	)
}

// NavbarLink is a link in the Navbar.
func NavbarLink(path, text string, active bool) Node {
	return A(Href(path), Text(text),
		// Apply CSS classes conditionally
		Classes{
			"px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:text-white focus:bg-gray-700": true,
			"text-white bg-gray-900":                           active,
			"text-gray-300 hover:text-white hover:bg-gray-700": !active,
		},
	)
}

func Container(children ...Node) Node {
	return Div(Class("max-w-7xl mx-auto px-2 sm:px-6 lg:px-8"), Group(children))
}

func Prose(children ...Node) Node {
	return Div(Class("prose"), Group(children))
}

func PageFooter() Node {
	return Footer(Class("prose prose-sm prose-indigo"),
		P(
			// We can use string interpolation directly, like fmt.Sprintf.
			Textf("Rendered %v. ", time.Now().Format(time.RFC3339)),

			// Conditional inclusion
			If(time.Now().Second()%2 == 0, Text("It's an even second.")),
			If(time.Now().Second()%2 == 1, Text("It's an odd second.")),
		),

		P(A(Href("https://www.gomponents.com"), Text("gomponents"))),
	)
}
