package page

import (
	"context"
	"testing"
)

func editPage(t *testing.T, html string) *Page {
	t.Helper()

	p, err := New(Config{HTML: html, Width: 320, Height: 240})
	if err != nil {
		t.Fatal(err)
	}

	if err := p.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	return p
}

func clickBox(t *testing.T, p *Page, id string) {
	t.Helper()

	for _, b := range p.Boxes() {
		if b.ID != id {
			continue
		}

		if err := p.Click(context.Background(), b.X+b.W/2, b.Y+b.H/2); err != nil {
			t.Fatal(err)
		}

		return
	}

	t.Fatalf("no box %q", id)
}

func TestBeforeEditFiresBeforeCut(t *testing.T) {
	ctx := context.Background()
	view := &struct{ Email string }{Email: "ab"}
	p := editPage(t, `<input id="e" type="text" data-bind="Email" style="display:block;width:120px;height:20px">`)
	p.SetData(view)

	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	clickBox(t, p, "e")

	var got Box
	p.Handle(Handlers{BeforeEdit: func(_ context.Context, b Box) error {
		got = b
		return nil
	}})

	text, ok, err := p.Cut(ctx)
	if err != nil || !ok || text != "ab" {
		t.Fatalf("cut = %q %v %v", text, ok, err)
	}

	if got.ID != "e" {
		t.Fatalf("before edit box = %+v", got)
	}

	if view.Email != "" {
		t.Fatalf("bound = %q", view.Email)
	}
}
