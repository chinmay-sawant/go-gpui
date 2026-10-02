package page

import (
	"context"
	"testing"
)

func TestBindE2EBackspaceWritesThroughToTemplate(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	view := bftView{Email: "abc"}
	p := bftPage(t, bftField)
	bftDraw(t, p, &view)
	bftClick(t, p, "email")
	if err := p.Backspace(ctx); err != nil {
		t.Fatal(err)
	}

	if view.Email != "ab" {
		t.Fatalf("view.Email %q", view.Email)
	}
	if p.FormValue("email") != "ab" {
		t.Fatalf("FormValue %q", p.FormValue("email"))
	}
	if got := bftBox(t, p, "out"); got != "ab" {
		t.Fatalf("paragraph %q", got)
	}
}

func TestBindE2EPasteWritesThroughToTemplate(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	view := bftView{}
	p := bftPage(t, bftField)
	bftDraw(t, p, &view)
	bftClick(t, p, "email")
	if err := p.Paste(ctx, "ada"); err != nil {
		t.Fatal(err)
	}

	if view.Email != "ada" {
		t.Fatalf("view.Email %q", view.Email)
	}
	if p.FormValue("email") != "ada" {
		t.Fatalf("FormValue %q", p.FormValue("email"))
	}
	if got := bftBox(t, p, "out"); got != "ada" {
		t.Fatalf("paragraph %q", got)
	}
}
