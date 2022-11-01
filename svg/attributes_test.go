package svg

import (
	"fmt"
	"testing"

	"github.com/melias122/html"
)

func TestSimpleAttributes(t *testing.T) {
	cases := map[string]func(string) html.Node{
		"clip-rule": ClipRule,
		"d":         D,
		"fill":      Fill,
		"fill-rule": FillRule,
		"stroke":    Stroke,
		"viewBox":   ViewBox,
	}

	for name, fn := range cases {
		t.Run(fmt.Sprintf(`should output %v="hat"`, name), func(t *testing.T) {
			n := html.El("element", fn("hat"))
			Equal(t, fmt.Sprintf(`<element %v="hat"></element>`, name), n)
		})
	}
}
