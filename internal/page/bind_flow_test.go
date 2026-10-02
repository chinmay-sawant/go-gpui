package page

import (
	"context"
	"strings"
	"testing"
)

type bftView struct {
	Email string
}

func bftPage(t *testing.T, body string) *Page {
	t.Helper()

	p, err := New(Config{
		HTML:   `<body style="margin:8px">` + body + `</body>`,
		Width:  640,
		Height: 480,
	})
	if err != nil {
		t.Fatal(err)
	}

	return p
}

func bftDraw(t *testing.T, p *Page, data any) {
	t.Helper()

	p.SetData(data)
	if err := p.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func bftBox(t *testing.T, p *Page, id string) string {
	t.Helper()

	for _, b := range p.Boxes() {
		if b.ID == id {
			return b.Text
		}
	}

	t.Fatalf("no box %q", id)

	return ""
}

func bftClick(t *testing.T, p *Page, id string) {
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

const bftField = `<input id="email" type="text" data-bind="Email" style="display:block;width:220px;height:28px"><p id="out">{{.Email}}</p>`

func TestBindE2ETypeWritesThroughToTemplate(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	view := bftView{}
	p := bftPage(t, bftField)
	bftDraw(t, p, &view)
	bftClick(t, p, "email")
	if err := p.Type(ctx, "ada"); err != nil {
		t.Fatal(err)
	}

	if view.Email != "ada" {
		t.Fatalf("view.Email %q", view.Email)
	}
	if p.FormValue("email") != "ada" {
		t.Fatalf("FormValue %q", p.FormValue("email"))
	}
	if got := bftBox(t, p, "out"); !strings.Contains(got, "ada") {
		t.Fatalf("paragraph %q", got)
	}
}
