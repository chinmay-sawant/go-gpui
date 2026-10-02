package page

import (
	"context"
	"testing"
)

type bdtView struct {
	Email string
	OK    bool
	Pick  string
}

const bdtField = `<input id="email" type="text" data-bind="Email" style="display:block;width:220px;height:28px"><p id="out">{{.Email}}</p>`

func TestBindE2ESetDataPushesChangedField(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	view := bdtView{Email: "ada"}
	p := bftPage(t, bdtField)
	bftDraw(t, p, &view)

	if p.FormValue("email") != "ada" || bftBox(t, p, "email") != "ada" {
		t.Fatalf("form %q box %q", p.FormValue("email"), bftBox(t, p, "email"))
	}

	view.Email = "grace"
	p.SetData(&view)
	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if p.FormValue("email") != "grace" {
		t.Fatalf("FormValue %q", p.FormValue("email"))
	}
	if got := bftBox(t, p, "out"); got != "grace" {
		t.Fatalf("paragraph %q", got)
	}
}

func TestBindE2EValueDataStaysInert(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	view := bdtView{Email: "grace"}
	p := bftPage(t, bdtField)
	bftDraw(t, p, view)

	if p.FormValue("email") != "" {
		t.Fatalf("FormValue %q", p.FormValue("email"))
	}

	bftClick(t, p, "email")
	if err := p.Type(ctx, "ada"); err != nil {
		t.Fatal(err)
	}

	if p.FormValue("email") != "ada" || view.Email != "grace" {
		t.Fatalf("FormValue %q view %q", p.FormValue("email"), view.Email)
	}
	if got := bftBox(t, p, "out"); got != "grace" {
		t.Fatalf("paragraph %q", got)
	}
}
