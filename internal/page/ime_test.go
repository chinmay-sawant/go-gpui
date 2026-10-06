package page

import (
	"context"
	"testing"
)

// imeField focuses a text field and types value into it.
func imeField(t *testing.T, value string) *Page {
	t.Helper()

	p := bedPage(t, `<input id="e" type="text"`+bedWide+`>`, nil)
	bedClick(t, p, "e")

	if value != "" {
		if err := p.Type(context.Background(), value); err != nil {
			t.Fatal(err)
		}
	}

	return p
}

func TestIMEControlNeedsFocus(t *testing.T) {
	t.Parallel()

	p := bedPage(t, `<input id="e" type="text"`+bedWide+`>`, nil)

	if _, _, _, _, ok := p.IMEContext(); ok {
		t.Fatal("context without focus")
	}
}

func TestIMEContextAroundCaret(t *testing.T) {
	t.Parallel()

	p := imeField(t, "hello")
	p.setCaret(2, false)

	box, caret, before, after, ok := p.IMEContext()
	if !ok || caret != 2 || before != "he" || after != "llo" {
		t.Fatalf("ok=%v caret=%d before=%q after=%q", ok, caret, before, after)
	}

	if box.W <= 0 || box.ID != "e" {
		t.Fatalf("box = %+v", box)
	}
}

func TestIMEContextCountsBytes(t *testing.T) {
	t.Parallel()

	p := imeField(t, "héllo")
	p.setCaret(2, false)

	_, caret, before, _, ok := p.IMEContext()
	if !ok || caret != 3 || before != "hé" {
		t.Fatalf("ok=%v caret=%d before=%q", ok, caret, before)
	}
}

func TestIMEContextFollowsTheScroll(t *testing.T) {
	t.Parallel()

	p := imeField(t, "hello")
	_, doc := 0, 0
	for _, b := range p.Boxes() {
		if b.ID == "e" {
			doc = int(b.Y)
		}
	}

	p.SetScrollOffset(0, 120)

	box, _, _, _, ok := p.IMEContext()
	if !ok {
		t.Fatal("no context")
	}

	if got := int(box.Y); got != doc-120 {
		t.Fatalf("box y = %d, want %d", got, doc-120)
	}
}
