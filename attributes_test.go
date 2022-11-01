package html

import (
	"fmt"
	"os"
	"testing"
)

func TestClasses(t *testing.T) {
	t.Run("given a map, returns sorted keys from the map with value true", func(t *testing.T) {
		Equal(t, ` class="boheme-hat hat partyhat"`, Classes{
			"boheme-hat": true,
			"hat":        true,
			"partyhat":   true,
			"turtlehat":  false,
		})
	})

	t.Run("renders as attribute in an element", func(t *testing.T) {
		e := El("div", Classes{"hat": true})
		Equal(t, `<div class="hat"></div>`, e)
	})

	t.Run("also works with fmt", func(t *testing.T) {
		a := Classes{"hat": true}
		if a.String() != ` class="hat"` {
			t.FailNow()
		}
	})
}

func ExampleClasses() {
	e := El("div", Classes{"party-hat": true, "boring-hat": false})
	_ = e.Render(os.Stdout)
	// Output: <div class="party-hat"></div>
}

func TestBooleanAttributes(t *testing.T) {
	cases := map[string]func() Node{
		"async":       Async,
		"autofocus":   AutoFocus,
		"autoplay":    AutoPlay,
		"controls":    Controls,
		"defer":       Defer,
		"disabled":    Disabled,
		"loop":        Loop,
		"multiple":    Multiple,
		"muted":       Muted,
		"playsinline": PlaysInline,
		"readonly":    ReadOnly,
		"required":    Required,
		"selected":    Selected,
	}

	for name, fn := range cases {
		t.Run(fmt.Sprintf("should output %v", name), func(t *testing.T) {
			n := El("div", fn())
			Equal(t, fmt.Sprintf(`<div %v></div>`, name), n)
		})
	}
}

func TestSimpleAttributes(t *testing.T) {
	cases := map[string]func(string) Node{
		"accept":       Accept,
		"action":       Action,
		"alt":          Alt,
		"as":           As,
		"autocomplete": AutoComplete,
		"charset":      Charset,
		"class":        Class,
		"cols":         Cols,
		"content":      Content,
		"enctype":      EncType,
		"for":          For,
		"form":         FormAttr,
		"height":       Height,
		"href":         Href,
		"id":           ID,
		"lang":         Lang,
		"loading":      Loading,
		"max":          Max,
		"maxlength":    MaxLength,
		"method":       Method,
		"min":          Min,
		"minlength":    MinLength,
		"name":         Name,
		"pattern":      Pattern,
		"placeholder":  Placeholder,
		"poster":       Poster,
		"preload":      Preload,
		"rel":          Rel,
		"role":         Role,
		"rows":         Rows,
		"src":          Src,
		"srcset":       SrcSet,
		"style":        StyleAttr,
		"tabindex":     TabIndex,
		"target":       Target,
		"title":        TitleAttr,
		"type":         Type,
		"value":        Value,
		"width":        Width,
	}

	for name, fn := range cases {
		t.Run(fmt.Sprintf(`should output %v="hat"`, name), func(t *testing.T) {
			n := El("div", fn("hat"))
			Equal(t, fmt.Sprintf(`<div %v="hat"></div>`, name), n)
		})
	}
}

func TestAria(t *testing.T) {
	t.Run("returns an attribute which name is prefixed with aria-", func(t *testing.T) {
		n := Aria("selected", "true")
		Equal(t, ` aria-selected="true"`, n)
	})
}

func TestDataAttr(t *testing.T) {
	t.Run("returns an attribute which name is prefixed with data-", func(t *testing.T) {
		n := DataAttr("id", "partyhat")
		Equal(t, ` data-id="partyhat"`, n)
	})
}
