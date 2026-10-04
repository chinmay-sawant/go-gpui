package page

import (
	"context"
	"testing"
)

func TestSelectAtOutsideBlurs(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p := bftPage(t, `<input id="t" type="text" value="ab"`+bedWide+`>`)
	bftDraw(t, p, nil)

	op := runOp(p, "ab")
	left, width := runSpan(p, op)
	base := runBase(p, op)

	if err := p.SelectAt(ctx, left+width/2, base-4); err != nil {
		t.Fatal(err)
	}

	if err := p.SelectAt(ctx, 1, 1); err != nil {
		t.Fatal(err)
	}

	if p.FocusID() != "" {
		t.Fatalf("focus %q after outside click", p.FocusID())
	}
}
