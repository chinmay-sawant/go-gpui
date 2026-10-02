package page

import (
	"context"
	"testing"
)

type bedView struct {
	Email string
	Agree bool
	Plan  string
}

const bedWide = ` style="display:block;width:220px;height:28px"`

func bedPage(t *testing.T, body string, data any) *Page {
	t.Helper()

	p, err := New(Config{HTML: `<body style="margin:8px">` + body + `</body>`, Width: 640, Height: 480})
	if err != nil {
		t.Fatal(err)
	}

	p.SetData(data)
	if err := p.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	return p
}

func bedClick(t *testing.T, p *Page, id string) {
	t.Helper()

	for _, b := range p.Boxes() {
		if b.ID == id {
			if err := p.Click(context.Background(), b.X+b.W/2, b.Y+b.H/2); err != nil {
				t.Fatal(err)
			}

			return
		}
	}

	t.Fatalf("no box %q", id)
}

func TestBindEditTypeWritesStruct(t *testing.T) {
	t.Parallel()

	view := &bedView{}
	p := bedPage(t, `<input id="e" type="text" data-bind="Email"`+bedWide+`>`, view)
	bedClick(t, p, "e")
	if err := p.Type(context.Background(), "ab"); err != nil {
		t.Fatal(err)
	}

	if view.Email != "ab" || p.FormValue("e") != "ab" {
		t.Fatalf("view %q form %q", view.Email, p.FormValue("e"))
	}
}

func TestBindEditUnboundWritesNothing(t *testing.T) {
	t.Parallel()

	view := &bedView{Email: "keep"}
	p := bedPage(t, `<input id="e" type="text"`+bedWide+`>`, view)
	bedClick(t, p, "e")
	if err := p.Type(context.Background(), "ab"); err != nil {
		t.Fatal(err)
	}

	if view.Email != "keep" {
		t.Fatalf("view %q", view.Email)
	}
}
