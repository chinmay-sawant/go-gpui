package page

import (
	"strings"
	"testing"
)

func TestRewriteSelectedNoCaret(t *testing.T) {
	t.Parallel()

	src := `<input id="e" type="text" value="x">`
	ctrl := Control{ID: "e", Tag: "input", Type: "text", Value: "x"}
	got := rewriteControls(src, []controlSpan{whole(src, ctrl)}, nil, "e", true)
	tag := openOf(got, "span")
	if !strings.Contains(tag, selectedAttr) || !strings.Contains(tag, focusAttr) {
		t.Fatalf("%s", got)
	}
	if strings.Contains(spanPart(got), "data-gpui-caret") {
		t.Fatalf("%s", got)
	}
}

func TestRewriteCaretOnly(t *testing.T) {
	t.Parallel()

	src := `<input id="e" type="text" value="x">`
	ctrl := Control{ID: "e", Tag: "input", Type: "text", Value: "x"}
	got := rewriteControls(src, []controlSpan{whole(src, ctrl)}, nil, "e", false)
	if !strings.Contains(got, `>x<span data-gpui-caret="1"></span></span>`) {
		t.Fatalf("%s", got)
	}
	got = rewriteControls(src, []controlSpan{whole(src, ctrl)}, nil, "", false)
	if strings.Contains(spanPart(got), "data-gpui-caret") || strings.Contains(spanPart(got), focusAttr) {
		t.Fatalf("%s", got)
	}
}
