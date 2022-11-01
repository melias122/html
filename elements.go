// Package html provides common HTML elements and attributes.
// See https://developer.mozilla.org/en-US/docs/Web/HTML/Element for a list of elements.
// See https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes for a list of attributes.
package html

import (
	"io"
)

// Doctype returns a special kind of Node that prefixes its sibling with the string "<!doctype html>".
func Doctype(sibling Node) Node {
	return NodeFunc(func(w io.Writer) error {
		if _, err := w.Write([]byte("<!doctype html>")); err != nil {
			return err
		}
		return sibling.Render(w)
	})
}

func A(children ...Node) Node {
	return El("a", children...)
}

func Address(children ...Node) Node {
	return El("address", children...)
}

func Area(children ...Node) Node {
	return El("area", children...)
}

func Article(children ...Node) Node {
	return El("article", children...)
}

func Aside(children ...Node) Node {
	return El("aside", children...)
}

func Audio(children ...Node) Node {
	return El("audio", children...)
}

func Base(children ...Node) Node {
	return El("base", children...)
}

func BlockQuote(children ...Node) Node {
	return El("blockquote", children...)
}

func Body(children ...Node) Node {
	return El("body", children...)
}

func Br(children ...Node) Node {
	return El("br", children...)
}

func Button(children ...Node) Node {
	return El("button", children...)
}

func Canvas(children ...Node) Node {
	return El("canvas", children...)
}

func Cite(children ...Node) Node {
	return El("cite", children...)
}

func Code(children ...Node) Node {
	return El("code", children...)
}

func Col(children ...Node) Node {
	return El("col", children...)
}

func ColGroup(children ...Node) Node {
	return El("colgroup", children...)
}

func DataEl(children ...Node) Node {
	return El("data", children...)
}

func DataList(children ...Node) Node {
	return El("datalist", children...)
}

func Details(children ...Node) Node {
	return El("details", children...)
}

func Dialog(children ...Node) Node {
	return El("dialog", children...)
}

func Div(children ...Node) Node {
	return El("div", children...)
}

func Dl(children ...Node) Node {
	return El("dl", children...)
}

func Embed(children ...Node) Node {
	return El("embed", children...)
}

func FormEl(children ...Node) Node {
	return El("form", children...)
}

func FieldSet(children ...Node) Node {
	return El("fieldset", children...)
}

func Figure(children ...Node) Node {
	return El("figure", children...)
}

func Footer(children ...Node) Node {
	return El("footer", children...)
}

func Head(children ...Node) Node {
	return El("head", children...)
}

func Header(children ...Node) Node {
	return El("header", children...)
}

func HGroup(children ...Node) Node {
	return El("hgroup", children...)
}

func Hr(children ...Node) Node {
	return El("hr", children...)
}

func HTML(children ...Node) Node {
	return El("html", children...)
}

func IFrame(children ...Node) Node {
	return El("iframe", children...)
}

func Img(children ...Node) Node {
	return El("img", children...)
}

func Input(children ...Node) Node {
	return El("input", children...)
}

func InputHidden(name, value string, children ...Node) Node {
	return Input(Type("hidden"), Name(name), Value(value), Group(children))
}

func Label(children ...Node) Node {
	return El("label", children...)
}

func Legend(children ...Node) Node {
	return El("legend", children...)
}

func Li(children ...Node) Node {
	return El("li", children...)
}

func Link(children ...Node) Node {
	return El("link", children...)
}

func LinkStylesheet(href string, children ...Node) Node {
	return Link(Rel("stylesheet"), Href(href), Group(children))
}

func LinkPreload(href, as string, children ...Node) Node {
	return Link(Rel("preload"), Href(href), As(as), Group(children))
}

func Main(children ...Node) Node {
	return El("main", children...)
}

func Menu(children ...Node) Node {
	return El("menu", children...)
}

func Meta(children ...Node) Node {
	return El("meta", children...)
}

func Meter(children ...Node) Node {
	return El("meter", children...)
}

func Nav(children ...Node) Node {
	return El("nav", children...)
}

func NoScript(children ...Node) Node {
	return El("noscript", children...)
}

func Object(children ...Node) Node {
	return El("object", children...)
}

func Ol(children ...Node) Node {
	return El("ol", children...)
}

func OptGroup(children ...Node) Node {
	return El("optgroup", children...)
}

func Option(children ...Node) Node {
	return El("option", children...)
}

func P(children ...Node) Node {
	return El("p", children...)
}

func Param(children ...Node) Node {
	return El("param", children...)
}

func Picture(children ...Node) Node {
	return El("picture", children...)
}

func Pre(children ...Node) Node {
	return El("pre", children...)
}

func Progress(children ...Node) Node {
	return El("progress", children...)
}

func Script(children ...Node) Node {
	return El("script", children...)
}

func Section(children ...Node) Node {
	return El("section", children...)
}

func Select(children ...Node) Node {
	return El("select", children...)
}

func Source(children ...Node) Node {
	return El("source", children...)
}

func Span(children ...Node) Node {
	return El("span", children...)
}

func StyleEl(children ...Node) Node {
	return El("style", children...)
}

func Summary(children ...Node) Node {
	return El("summary", children...)
}

func SVG(children ...Node) Node {
	return El("svg", children...)
}

func Table(children ...Node) Node {
	return El("table", children...)
}

func TBody(children ...Node) Node {
	return El("tbody", children...)
}

func Td(children ...Node) Node {
	return El("td", children...)
}

func Textarea(children ...Node) Node {
	return El("textarea", children...)
}

func TFoot(children ...Node) Node {
	return El("tfoot", children...)
}

func Th(children ...Node) Node {
	return El("th", children...)
}

func THead(children ...Node) Node {
	return El("thead", children...)
}

func Tr(children ...Node) Node {
	return El("tr", children...)
}

func Ul(children ...Node) Node {
	return El("ul", children...)
}

func Wbr(children ...Node) Node {
	return El("wbr", children...)
}

func Abbr(children ...Node) Node {
	return El("abbr", Group(children))
}

func B(children ...Node) Node {
	return El("b", Group(children))
}

func Caption(children ...Node) Node {
	return El("caption", Group(children))
}

func Dd(children ...Node) Node {
	return El("dd", Group(children))
}

func Del(children ...Node) Node {
	return El("del", Group(children))
}

func Dfn(children ...Node) Node {
	return El("dfn", Group(children))
}

func Dt(children ...Node) Node {
	return El("dt", Group(children))
}

func Em(children ...Node) Node {
	return El("em", Group(children))
}

func FigCaption(children ...Node) Node {
	return El("figcaption", Group(children))
}

func H1(children ...Node) Node {
	return El("h1", Group(children))
}

func H2(children ...Node) Node {
	return El("h2", Group(children))
}

func H3(children ...Node) Node {
	return El("h3", Group(children))
}

func H4(children ...Node) Node {
	return El("h4", Group(children))
}

func H5(children ...Node) Node {
	return El("h5", Group(children))
}

func H6(children ...Node) Node {
	return El("h6", Group(children))
}

func I(children ...Node) Node {
	return El("i", Group(children))
}

func Ins(children ...Node) Node {
	return El("ins", Group(children))
}

func Kbd(children ...Node) Node {
	return El("kbd", Group(children))
}

func Mark(children ...Node) Node {
	return El("mark", Group(children))
}

func Q(children ...Node) Node {
	return El("q", Group(children))
}

func S(children ...Node) Node {
	return El("s", Group(children))
}

func Samp(children ...Node) Node {
	return El("samp", Group(children))
}

func Small(children ...Node) Node {
	return El("small", Group(children))
}

func Strong(children ...Node) Node {
	return El("strong", Group(children))
}

func Sub(children ...Node) Node {
	return El("sub", Group(children))
}

func Sup(children ...Node) Node {
	return El("sup", Group(children))
}

func Time(children ...Node) Node {
	return El("time", Group(children))
}

func TitleEl(children ...Node) Node {
	return El("title", Group(children))
}

func U(children ...Node) Node {
	return El("u", Group(children))
}

func Var(children ...Node) Node {
	return El("var", Group(children))
}

func Video(children ...Node) Node {
	return El("video", Group(children))
}
