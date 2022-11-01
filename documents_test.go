package html

import (
	"testing"
)

func TestHTML5(t *testing.T) {
	t.Run("returns an html5 document template", func(t *testing.T) {
		e := HTML5(HTML5Props{
			Title:       "Hat",
			Description: "Love hats.",
			Language:    "en",
			Head:        []Node{Link(Rel("stylesheet"), Href("/hat.css"))},
			Body:        []Node{Div()},
		})

		Equal(t, `<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><title>Hat</title><meta name="description" content="Love hats."><link rel="stylesheet" href="/hat.css"></head><body><div></div></body></html>`, e)
	})

	t.Run("returns no language, description, and extra head/body elements if empty", func(t *testing.T) {
		e := HTML5(HTML5Props{
			Title: "Hat",
		})

		Equal(t, `<!doctype html><html><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><title>Hat</title></head><body></body></html>`, e)
	})
}
