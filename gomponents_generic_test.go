//go:build go1.18
// +build go1.18

package html

import (
	"os"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("maps slices to nodes", func(t *testing.T) {
		items := []string{"hat", "partyhat", "turtlehat"}
		lis := Map(items, func(i string) Node {
			return El("li", Text(i))
		})

		list := El("ul", lis...)

		Equal(t, `<ul><li>hat</li><li>partyhat</li><li>turtlehat</li></ul>`, list)
	})
}

func ExampleMap() {
	items := []string{"party hat", "super hat"}
	e := El("ul", Group(Map(items, func(i string) Node {
		return El("li", Text(i))
	})))
	_ = e.Render(os.Stdout)
	// Output: <ul><li>party hat</li><li>super hat</li></ul>
}
