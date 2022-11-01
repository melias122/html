package html

import (
	"errors"
	"fmt"
	"testing"
)

type erroringWriter struct{}

func (w *erroringWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("don't want to write")
}

func TestDoctype(t *testing.T) {
	t.Run("returns doctype and children", func(t *testing.T) {
		Equal(t, `<!doctype html><html></html>`, Doctype(El("html")))
	})

	t.Run("errors on write error in Render", func(t *testing.T) {
		err := Doctype(El("html")).Render(&erroringWriter{})
		Error(t, err)
	})
}

func TestSimpleElements(t *testing.T) {
	cases := map[string]func(...Node) Node{
		"a":          A,
		"abbr":       Abbr,
		"address":    Address,
		"article":    Article,
		"aside":      Aside,
		"audio":      Audio,
		"b":          B,
		"blockquote": BlockQuote,
		"body":       Body,
		"button":     Button,
		"canvas":     Canvas,
		"caption":    Caption,
		"cite":       Cite,
		"code":       Code,
		"colgroup":   ColGroup,
		"data":       DataEl,
		"datalist":   DataList,
		"dd":         Dd,
		"del":        Del,
		"details":    Details,
		"dfn":        Dfn,
		"dialog":     Dialog,
		"div":        Div,
		"dl":         Dl,
		"dt":         Dt,
		"em":         Em,
		"fieldset":   FieldSet,
		"figcaption": FigCaption,
		"figure":     Figure,
		"footer":     Footer,
		"form":       FormEl,
		"h1":         H1,
		"h2":         H2,
		"h3":         H3,
		"h4":         H4,
		"h5":         H5,
		"h6":         H6,
		"head":       Head,
		"header":     Header,
		"hgroup":     HGroup,
		"html":       HTML,
		"i":          I,
		"iframe":     IFrame,
		"ins":        Ins,
		"kbd":        Kbd,
		"label":      Label,
		"legend":     Legend,
		"li":         Li,
		"main":       Main,
		"mark":       Mark,
		"menu":       Menu,
		"meter":      Meter,
		"nav":        Nav,
		"noscript":   NoScript,
		"object":     Object,
		"ol":         Ol,
		"optgroup":   OptGroup,
		"option":     Option,
		"p":          P,
		"picture":    Picture,
		"pre":        Pre,
		"progress":   Progress,
		"q":          Q,
		"s":          S,
		"samp":       Samp,
		"script":     Script,
		"section":    Section,
		"select":     Select,
		"small":      Small,
		"span":       Span,
		"strong":     Strong,
		"style":      StyleEl,
		"sub":        Sub,
		"summary":    Summary,
		"sup":        Sup,
		"svg":        SVG,
		"table":      Table,
		"tbody":      TBody,
		"td":         Td,
		"textarea":   Textarea,
		"tfoot":      TFoot,
		"th":         Th,
		"thead":      THead,
		"time":       Time,
		"title":      TitleEl,
		"tr":         Tr,
		"u":          U,
		"ul":         Ul,
		"var":        Var,
		"video":      Video,
	}

	for name, fn := range cases {
		t.Run(fmt.Sprintf("should output %v", name), func(t *testing.T) {
			n := fn(Attr("id", "hat"))
			Equal(t, fmt.Sprintf(`<%v id="hat"></%v>`, name, name), n)
		})
	}
}

func TestSimpleVoidKindElements(t *testing.T) {
	cases := map[string]func(...Node) Node{
		"area":   Area,
		"base":   Base,
		"br":     Br,
		"col":    Col,
		"embed":  Embed,
		"hr":     Hr,
		"img":    Img,
		"input":  Input,
		"link":   Link,
		"meta":   Meta,
		"param":  Param,
		"source": Source,
		"wbr":    Wbr,
	}

	for name, fn := range cases {
		t.Run(fmt.Sprintf("should output %v", name), func(t *testing.T) {
			n := fn(Attr("id", "hat"))
			Equal(t, fmt.Sprintf(`<%v id="hat">`, name), n)
		})
	}
}

func TestInputHidden(t *testing.T) {
	t.Run("returns an input element with type hidden, and the given name and value", func(t *testing.T) {
		n := InputHidden("id", "partyhat", Attr("class", "hat"))
		Equal(t, `<input type="hidden" name="id" value="partyhat" class="hat">`, n)
	})
}

func TestLinkStylesheet(t *testing.T) {
	t.Run("returns a link element with rel stylesheet and the given href", func(t *testing.T) {
		n := LinkStylesheet("style.css", Attr("media", "print"))
		Equal(t, `<link rel="stylesheet" href="style.css" media="print">`, n)
	})
}

func TestLinkPreload(t *testing.T) {
	t.Run("returns a link element with rel preload and the given href and as", func(t *testing.T) {
		n := LinkPreload("party.woff2", "font", Attr("type", "font/woff2"))
		Equal(t, `<link rel="preload" href="party.woff2" as="font" type="font/woff2">`, n)
	})
}
