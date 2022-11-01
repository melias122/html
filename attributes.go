package html

import (
	"io"
	"sort"
	"strings"
)

// Classes is a map of strings to booleans, which Renders to an attribute with name "class".
// The attribute value is a sorted, space-separated string of all the map keys,
// for which the corresponding map value is true.
type Classes map[string]bool

func (c Classes) Render(w io.Writer) error {
	var included []string
	for c, include := range c {
		if include {
			included = append(included, c)
		}
	}
	sort.Strings(included)
	return Class(strings.Join(included, " ")).Render(w)
}

func (c Classes) Type() NodeType {
	return AttributeType
}

// String satisfies fmt.Stringer.
func (c Classes) String() string {
	var b strings.Builder
	_ = c.Render(&b)
	return b.String()
}

func Async() Node {
	return Attr("async")
}

func AutoFocus() Node {
	return Attr("autofocus")
}

func AutoPlay() Node {
	return Attr("autoplay")
}

func Controls() Node {
	return Attr("controls")
}

func Defer() Node {
	return Attr("defer")
}

func Disabled() Node {
	return Attr("disabled")
}

func Loop() Node {
	return Attr("loop")
}

func Multiple() Node {
	return Attr("multiple")
}

func Muted() Node {
	return Attr("muted")
}

func PlaysInline() Node {
	return Attr("playsinline")
}

func ReadOnly() Node {
	return Attr("readonly")
}

func Required() Node {
	return Attr("required")
}

func Selected() Node {
	return Attr("selected")
}

func Accept(v string) Node {
	return Attr("accept", v)
}

func Action(v string) Node {
	return Attr("action", v)
}

func Alt(v string) Node {
	return Attr("alt", v)
}

// Aria attributes automatically have their name prefixed with "aria-".
func Aria(name, v string) Node {
	return Attr("aria-"+name, v)
}

func As(v string) Node {
	return Attr("as", v)
}

func AutoComplete(v string) Node {
	return Attr("autocomplete", v)
}

func Charset(v string) Node {
	return Attr("charset", v)
}

func Class(v string) Node {
	return Attr("class", v)
}

func Cols(v string) Node {
	return Attr("cols", v)
}

func Content(v string) Node {
	return Attr("content", v)
}

// DataAttr attributes automatically have their name prefixed with "data-".
func DataAttr(name, v string) Node {
	return Attr("data-"+name, v)
}

func For(v string) Node {
	return Attr("for", v)
}

func FormAttr(v string) Node {
	return Attr("form", v)
}

func Height(v string) Node {
	return Attr("height", v)
}

func Href(v string) Node {
	return Attr("href", v)
}

func ID(v string) Node {
	return Attr("id", v)
}

func Lang(v string) Node {
	return Attr("lang", v)
}

func Loading(v string) Node {
	return Attr("loading", v)
}

func Max(v string) Node {
	return Attr("max", v)
}

func MaxLength(v string) Node {
	return Attr("maxlength", v)
}

func Method(v string) Node {
	return Attr("method", v)
}

func Min(v string) Node {
	return Attr("min", v)
}

func MinLength(v string) Node {
	return Attr("minlength", v)
}

func Name(v string) Node {
	return Attr("name", v)
}

func Pattern(v string) Node {
	return Attr("pattern", v)
}

func Placeholder(v string) Node {
	return Attr("placeholder", v)
}

func Poster(v string) Node {
	return Attr("poster", v)
}

func Preload(v string) Node {
	return Attr("preload", v)
}

func Rel(v string) Node {
	return Attr("rel", v)
}

func Role(v string) Node {
	return Attr("role", v)
}

func Rows(v string) Node {
	return Attr("rows", v)
}

func Src(v string) Node {
	return Attr("src", v)
}

func SrcSet(v string) Node {
	return Attr("srcset", v)
}

func StyleAttr(v string) Node {
	return Attr("style", v)
}

func TabIndex(v string) Node {
	return Attr("tabindex", v)
}

func Target(v string) Node {
	return Attr("target", v)
}

func TitleAttr(v string) Node {
	return Attr("title", v)
}

func Type(v string) Node {
	return Attr("type", v)
}

func Value(v string) Node {
	return Attr("value", v)
}

func Width(v string) Node {
	return Attr("width", v)
}

func EncType(v string) Node {
	return Attr("enctype", v)
}
