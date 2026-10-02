package page

import (
	"strings"
	"testing"
)

func TestRewriteTextareaFocus(t *testing.T) {
	t.Parallel()

	src := `<textarea id="note" name="n">old</textarea>`
	ctrl := Control{ID: "note", Tag: "textarea", Name: "n", Value: `a<b&c"`}
	got := rewriteControls(src, []controlSpan{whole(src, ctrl)}, nil, "note")
	tag := openOf(got, "textarea")
	if !strings.Contains(tag, `data-gpui-focus="1"`) || !strings.Contains(tag, `name="n"`) {
		t.Fatalf("%s", got)
	}
	if !strings.Contains(got, ">a&lt;b&amp;c&#34;</textarea>") {
		t.Fatalf("%s", got)
	}
}
