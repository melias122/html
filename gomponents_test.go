package html

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// Equal checks for equality between the given expected string and the rendered Node string.
func Equal(t *testing.T, expected string, actual Node) {
	t.Helper()

	var b strings.Builder
	_ = actual.Render(&b)
	if expected != b.String() {
		t.Fatalf(`expected "%v" but got "%v"`, expected, b.String())
	}
}

// Error checks for a non-nil error.
func Error(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Fatal("error is nil")
	}
}

func TestNodeFunc(t *testing.T) {
	t.Run("implements fmt.Stringer", func(t *testing.T) {
		fn := NodeFunc(func(w io.Writer) error {
			_, _ = w.Write([]byte("hat"))
			return nil
		})
		if fn.String() != "hat" {
			t.FailNow()
		}
	})
}

func TestAttr(t *testing.T) {
	t.Run("renders just the local name with one argument", func(t *testing.T) {
		a := Attr("required")
		Equal(t, " required", a)
	})

	t.Run("renders the name and value when given two arguments", func(t *testing.T) {
		a := Attr("id", "hat")
		Equal(t, ` id="hat"`, a)
	})

	t.Run("panics with more than two arguments", func(t *testing.T) {
		called := false
		defer func() {
			if err := recover(); err != nil {
				called = true
			}
		}()
		Attr("name", "value", "what is this?")
		if !called {
			t.FailNow()
		}
	})

	t.Run("implements fmt.Stringer", func(t *testing.T) {
		a := Attr("required")
		s := fmt.Sprintf("%v", a)
		if s != " required" {
			t.FailNow()
		}
	})

	t.Run("escapes attribute values", func(t *testing.T) {
		a := Attr(`id`, `hat"><script`)
		Equal(t, ` id="hat&#34;&gt;&lt;script"`, a)
	})
}

func BenchmarkAttr(b *testing.B) {
	b.Run("boolean attributes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := Attr("hat")
			_ = a.Render(&strings.Builder{})
		}
	})

	b.Run("name-value attributes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			a := Attr("hat", "party")
			_ = a.Render(&strings.Builder{})
		}
	})
}

func ExampleAttr_bool() {
	e := El("input", Attr("required"))
	_ = e.Render(os.Stdout)
	// Output: <input required>
}

func ExampleAttr_name_value() {
	e := El("div", Attr("id", "hat"))
	_ = e.Render(os.Stdout)
	// Output: <div id="hat"></div>
}

type outsider struct{}

func (o outsider) String() string {
	return "outsider"
}

func (o outsider) Render(w io.Writer) error {
	_, _ = w.Write([]byte("outsider"))
	return nil
}

func TestEl(t *testing.T) {
	t.Run("renders an empty element if no children given", func(t *testing.T) {
		e := El("div")
		Equal(t, "<div></div>", e)
	})

	t.Run("renders an empty element without closing tag if it's a void kind element", func(t *testing.T) {
		e := El("hr")
		Equal(t, "<hr>", e)

		e = El("br")
		Equal(t, "<br>", e)

		e = El("img")
		Equal(t, "<img>", e)
	})

	t.Run("renders an empty element if only attributes given as children", func(t *testing.T) {
		e := El("div", Attr("class", "hat"))
		Equal(t, `<div class="hat"></div>`, e)
	})

	t.Run("renders an element, attributes, and element children", func(t *testing.T) {
		e := El("div", Attr("class", "hat"), El("br"))
		Equal(t, `<div class="hat"><br></div>`, e)
	})

	t.Run("renders attributes at the correct place regardless of placement in parameter list", func(t *testing.T) {
		e := El("div", El("br"), Attr("class", "hat"))
		Equal(t, `<div class="hat"><br></div>`, e)
	})

	t.Run("renders outside if node does not implement nodeTypeDescriber", func(t *testing.T) {
		e := El("div", outsider{})
		Equal(t, `<div>outsider</div>`, e)
	})

	t.Run("does not fail on nil node", func(t *testing.T) {
		e := El("div", nil, El("br"), nil, El("br"))
		Equal(t, `<div><br><br></div>`, e)
	})

	t.Run("returns render error on cannot write", func(t *testing.T) {
		e := El("div")
		err := e.Render(&erroringWriter{})
		Error(t, err)
	})
}

func BenchmarkEl(b *testing.B) {
	b.Run("normal elements", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			e := El("div")
			_ = e.Render(&strings.Builder{})
		}
	})
}

func ExampleEl() {
	e := El("div", El("span"))
	_ = e.Render(os.Stdout)
	// Output: <div><span></span></div>
}

func TestText(t *testing.T) {
	t.Run("renders escaped text", func(t *testing.T) {
		e := Text("<div>")
		Equal(t, "&lt;div&gt;", e)
	})
}

func ExampleText() {
	e := El("span", Text("Party hats > normal hats."))
	_ = e.Render(os.Stdout)
	// Output: <span>Party hats &gt; normal hats.</span>
}

func TestTextf(t *testing.T) {
	t.Run("renders interpolated and escaped text", func(t *testing.T) {
		e := Textf("<%v>", "div")
		Equal(t, "&lt;div&gt;", e)
	})
}

func ExampleTextf() {
	e := El("span", Textf("%v party hats > %v normal hats.", 2, 3))
	_ = e.Render(os.Stdout)
	// Output: <span>2 party hats &gt; 3 normal hats.</span>
}

func TestRaw(t *testing.T) {
	t.Run("renders raw text", func(t *testing.T) {
		e := Raw("<div>")
		Equal(t, "<div>", e)
	})
}

func ExampleRaw() {
	e := El("span",
		Raw(`<button onclick="javascript:alert('Party time!')">Party hats</button> &gt; normal hats.`),
	)
	_ = e.Render(os.Stdout)
	// Output: <span><button onclick="javascript:alert('Party time!')">Party hats</button> &gt; normal hats.</span>
}

func TestGroup(t *testing.T) {
	t.Run("groups multiple nodes into one", func(t *testing.T) {
		children := []Node{El("br", Attr("id", "hat")), El("hr")}
		e := El("div", Attr("class", "foo"), El("img"), Group(children))
		Equal(t, `<div class="foo"><img><br id="hat"><hr></div>`, e)
	})

	t.Run("panics on direct render", func(t *testing.T) {
		e := Group(nil)
		panicked := false
		defer func() {
			if err := recover(); err != nil {
				panicked = true
			}
		}()
		_ = e.Render(nil)
		if !panicked {
			t.FailNow()
		}
	})

	t.Run("panics on direct string", func(t *testing.T) {
		e := Group(nil).(fmt.Stringer)
		panicked := false
		defer func() {
			if err := recover(); err != nil {
				panicked = true
			}
		}()
		_ = e.String()
		if !panicked {
			t.FailNow()
		}
	})
}

func TestIf(t *testing.T) {
	t.Run("returns node if condition is true", func(t *testing.T) {
		n := El("div", If(true, El("span")))
		Equal(t, "<div><span></span></div>", n)
	})

	t.Run("returns nil if condition is false", func(t *testing.T) {
		n := El("div", If(false, El("span")))
		Equal(t, "<div></div>", n)
	})
}

func ExampleIf() {
	showMessage := true
	e := El("div",
		If(showMessage, El("span", Text("You lost your hat!"))),
		If(!showMessage, El("span", Text("No messages."))),
	)
	_ = e.Render(os.Stdout)
	// Output: <div><span>You lost your hat!</span></div>
}
