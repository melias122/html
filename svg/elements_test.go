package svg

import (
	"fmt"
	"strings"
	"testing"

	"github.com/melias122/html"
)

// Equal checks for equality between the given expected string and the rendered Node string.
func Equal(t *testing.T, expected string, actual html.Node) {
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

func TestSimpleElements(t *testing.T) {
	cases := map[string]func(...html.Node) html.Node{
		"path": Path,
	}

	for name, fn := range cases {
		t.Run(fmt.Sprintf("should output %v", name), func(t *testing.T) {
			n := fn(html.Attr("id", "hat"))
			Equal(t, fmt.Sprintf(`<%v id="hat"></%v>`, name, name), n)
		})
	}
}

func TestSVG(t *testing.T) {
	t.Run("outputs svg element with xml namespace attribute", func(t *testing.T) {
		Equal(t, `<svg xmlns="http://www.w3.org/2000/svg"><path></path></svg>`, SVG(html.El("path")))
	})
}
