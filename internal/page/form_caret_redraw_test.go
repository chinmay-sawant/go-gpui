package page

import (
	"context"
	"testing"
)

func TestCaretSurvivesRedraw(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p := bftPage(t, `<input id="t" type="text" value="abcdef"`+bedWide+`>`)
	bftDraw(t, p, nil)
	bftClick(t, p, "t")
	p.setCaret(4, false)

	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if p.form.caret != 4 || p.form.anchor != 4 || p.FocusID() != "t" {
		t.Fatalf("caret %d anchor %d focus %q", p.form.caret, p.form.anchor, p.FocusID())
	}
}
