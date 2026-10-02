package page

import (
	"context"
	"testing"
)

func TestFormSelectedTracksSelectAll(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	body := `<input id="a" type="text"` + bedWide + `><input id="b" type="text"` + bedWide + `>`
	p := bftPage(t, body)
	bftDraw(t, p, nil)
	bftClick(t, p, "a")
	if err := p.Type(ctx, "ab"); err != nil {
		t.Fatal(err)
	}
	if p.FormSelected("a") {
		t.Fatal("selected before SelectAll")
	}
	if err := p.SelectAll(ctx); err != nil {
		t.Fatal(err)
	}

	if !p.FormSelected("a") || p.FormSelected("b") {
		t.Fatalf("after SelectAll a=%v b=%v", p.FormSelected("a"), p.FormSelected("b"))
	}

	if err := p.Type(ctx, "c"); err != nil {
		t.Fatal(err)
	}

	if p.FormSelected("a") {
		t.Fatal("selected after Type")
	}
}

func TestFormSelectedFalseAfterBlur(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p := bftPage(t, `<input id="a" type="text"`+bedWide+`><div id="plain">x</div>`)
	bftDraw(t, p, nil)
	bftClick(t, p, "a")
	if err := p.SelectAll(ctx); err != nil {
		t.Fatal(err)
	}
	if !p.FormSelected("a") {
		t.Fatal("not selected before blur")
	}

	bftClick(t, p, "plain")

	if p.FormSelected("a") {
		t.Fatal("selected after blur")
	}
}

func TestFormSelectedNilForm(t *testing.T) {
	t.Parallel()

	p := &Page{}
	if p.FormSelected("a") {
		t.Fatal("nil form reports selected")
	}
}
