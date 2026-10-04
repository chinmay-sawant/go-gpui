package page

import (
	"context"
	"testing"
)

func focusFields(t *testing.T) *Page {
	t.Helper()

	p := bftPage(t, `<input id="a" type="text"`+bedWide+`>`+
		`<input id="b" type="text" disabled`+bedWide+`>`+
		`<input id="c" type="text" tabindex="-1"`+bedWide+`>`+
		`<input id="d" type="text"`+bedWide+`>`)
	bftDraw(t, p, nil)

	return p
}

func TestFocusTraversal(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p := focusFields(t)

	steps := []struct {
		next func(context.Context) error
		want string
	}{
		{p.FocusNext, "a"},
		{p.FocusNext, "d"},
		{p.FocusNext, "a"},
		{p.FocusPrev, "d"},
		{p.FocusPrev, "a"},
	}

	for i, step := range steps {
		if err := step.next(ctx); err != nil {
			t.Fatal(err)
		}

		if got := p.FocusID(); got != step.want {
			t.Fatalf("step %d focus %q want %q", i, got, step.want)
		}
	}
}

func TestFocusSkipsDisabled(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p := focusFields(t)

	if err := p.Focus(ctx, "b"); err != nil {
		t.Fatal(err)
	}

	if p.FocusID() != "" {
		t.Fatalf("disabled focus %q", p.FocusID())
	}

	if err := p.Focus(ctx, "c"); err != nil {
		t.Fatal(err)
	}

	if p.FocusID() != "c" {
		t.Fatalf("programmatic focus %q", p.FocusID())
	}

	if err := p.FocusNext(ctx); err != nil {
		t.Fatal(err)
	}

	if p.FocusID() != "a" {
		t.Fatalf("next from tabindex=-1 is %q", p.FocusID())
	}
}
