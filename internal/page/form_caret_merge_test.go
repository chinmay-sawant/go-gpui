package page

import (
	"context"
	"testing"
)

func TestCaretReplacesSelection(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p := bftPage(t, `<input id="t" type="text" value="abcd"`+bedWide+`>`)
	bftDraw(t, p, nil)
	bftClick(t, p, "t")
	p.form.anchor, p.form.caret = 1, 3

	if err := p.Type(ctx, "Z"); err != nil {
		t.Fatal(err)
	}

	if p.FormValue("t") != "aZd" || p.form.caret != 2 || p.form.anchor != 2 {
		t.Fatalf("value %q caret %d anchor %d", p.FormValue("t"), p.form.caret, p.form.anchor)
	}
}

func TestSelectAllSetsFullRange(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p := bftPage(t, `<input id="t" type="text" value="abcd"`+bedWide+`>`)
	bftDraw(t, p, nil)
	bftClick(t, p, "t")

	if err := p.SelectAll(ctx); err != nil {
		t.Fatal(err)
	}

	start, end := p.form.bounds(4)
	if start != 0 || end != 4 || p.form.caret != 4 || !p.form.all {
		t.Fatalf("range [%d,%d) caret %d all %v", start, end, p.form.caret, p.form.all)
	}

	if !p.FormSelected("t") {
		t.Fatal("not selected")
	}
}
