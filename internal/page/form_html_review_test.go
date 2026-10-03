package page

import (
	"strings"
	"testing"
)

func TestHeadOpenQuote(t *testing.T) {
	t.Parallel()

	src := `<header><head data-x="a>b"><body>`
	want := strings.Index(src, "<body>")
	if got := headOpen(src); got != want {
		t.Fatalf("headOpen: %d != %d", got, want)
	}
}

func TestHeadCSSBoundaries(t *testing.T) {
	t.Parallel()

	head := "<head>"
	input := `<input id="e">`
	ctrl := Control{ID: "e", Tag: "input", Type: "text", Value: "v"}
	src := head + input
	sp := controlSpan{Control: ctrl, Start: len(head), End: len(src)}
	if got := rewriteControls(src, []controlSpan{sp}, nil, "", false); !strings.HasPrefix(got, head+formCSS) {
		t.Fatalf("head boundary: %s", got)
	}

	src = input + head
	sp = controlSpan{Control: ctrl, Start: 0, End: len(input)}
	if got := rewriteControls(src, []controlSpan{sp}, nil, "", false); !strings.HasSuffix(got, head+formCSS) {
		t.Fatalf("source end: %s", got)
	}
}

func TestAttrValueQuotes(t *testing.T) {
	t.Parallel()

	cases := []struct{ raw, want string }{
		{`<input placeholder="'A &amp; B'">`, `'A & B'`},
		{`<input placeholder=A&amp;B>`, "A&B"},
	}
	for _, tc := range cases {
		if got := attrValue(tc.raw, "placeholder"); got != tc.want {
			t.Errorf("%q != %q", got, tc.want)
		}
	}
}

func TestTextareaCSSMarker(t *testing.T) {
	t.Parallel()

	src := `<textarea id="note"></textarea>`
	ctrl := Control{ID: "note", Tag: "textarea"}
	sp := controlSpan{Control: ctrl, Start: 0, End: len(src)}
	got := rewriteControls(src, []controlSpan{sp}, nil, "", false)
	if !strings.Contains(got, `data-gpui-field="textarea"`) ||
		!strings.Contains(formCSS, `white-space:pre-wrap`) {
		t.Fatalf("CSS mismatch: %s", got)
	}
}
