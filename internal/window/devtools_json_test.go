package window

import (
	"encoding/json"
	"image/color"
	"strings"
	"testing"
)

type devJSONTestRect struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type devJSONTestDoc struct {
	Tag   string            `json:"tag"`
	ID    string            `json:"id,omitempty"`
	Skip  string            `json:"-"`
	Items []devJSONTestRect `json:"items"`
	Rect  devJSONTestRect   `json:"rect"`
	Ops   int               `json:"ops"`
	OK    bool              `json:"ok"`
}

func devJSONFixture() devJSONTestDoc {
	return devJSONTestDoc{
		Tag:   "div",
		Skip:  "no",
		Items: []devJSONTestRect{{X: 1, Y: 2}, {X: 3, Y: 4}},
		Rect:  devJSONTestRect{X: 5, Y: 6},
		Ops:   2,
		OK:    true,
	}
}

// TestDevJSONValid checks the joined text parses as JSON.
func TestDevJSONValid(t *testing.T) {
	t.Parallel()

	text := devRowsText(devJSONRows(devJSONFixture(), nil))
	if !json.Valid([]byte(text)) {
		t.Fatalf("json = %q", text)
	}
}

// TestDevJSONOrderAndOmit checks declaration order and omitempty.
func TestDevJSONOrderAndOmit(t *testing.T) {
	t.Parallel()

	text := devRowsText(devJSONRows(devJSONFixture(), nil))
	if strings.Index(text, `"tag"`) > strings.Index(text, `"rect"`) {
		t.Fatalf("tag after rect: %q", text)
	}

	if strings.Contains(text, `"id"`) || strings.Contains(text, `"Skip"`) {
		t.Fatalf("omitempty or json:- leaked: %q", text)
	}
}

// TestDevJSONColors checks the syntax palette.
func TestDevJSONColors(t *testing.T) {
	t.Parallel()

	want := map[string]color.RGBA{
		"tag":   devKeyInk,
		`"div"`: devStrInk,
		"2":     devNumInk,
		"true":  devBoolInk,
	}

	for _, row := range devJSONRows(devJSONFixture(), nil) {
		for _, span := range row.line.spans {
			if ink, ok := want[span.text]; ok && span.ink != ink {
				t.Fatalf("span %q ink = %v, want %v", span.text, span.ink, ink)
			}
		}
	}
}
