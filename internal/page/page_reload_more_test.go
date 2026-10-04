package page_test

import (
	"context"
	"testing"
)

func TestReloadKeepsTypedField(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	path := writeSource(t, "index.html", `<input id="note" type="text"><p id="p">before</p>`)
	p := watchedPage(t, path)

	x, y := boxCenter(t, p, "note")
	if err := p.Click(ctx, x, y); err != nil {
		t.Fatal(err)
	}

	if err := p.Type(ctx, "hello"); err != nil {
		t.Fatal(err)
	}

	gen := p.Generation()
	next := `<input id="note" type="text"><p id="p">after</p><p id="extra">an extra paragraph</p>`
	rewriteSource(t, path, next)

	changed, err := p.PollReload(ctx)
	if !changed || err != nil {
		t.Fatalf("poll = %v, %v", changed, err)
	}

	if p.FormValue("note") != "hello" {
		t.Fatalf("value = %q", p.FormValue("note"))
	}

	if p.FocusedField() != "note" {
		t.Fatalf("focus = %q", p.FocusedField())
	}

	if p.Generation() != gen+1 {
		t.Fatalf("generation moved by %d", p.Generation()-gen)
	}
}

func TestReloadDropsRemovedFieldFocus(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	path := writeSource(t, "index.html", `<input id="note" type="text">`)
	p := watchedPage(t, path)

	x, y := boxCenter(t, p, "note")
	if err := p.Click(ctx, x, y); err != nil {
		t.Fatal(err)
	}

	rewriteSource(t, path, `<p id="p">no field</p>`)

	changed, err := p.PollReload(ctx)
	if !changed || err != nil {
		t.Fatalf("poll = %v, %v", changed, err)
	}

	if p.FocusedField() != "" {
		t.Fatalf("focus = %q, want none", p.FocusedField())
	}
}
