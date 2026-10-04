package page_test

import (
	"context"
	"testing"
)

func TestIdenticalRedrawNoDirty(t *testing.T) {
	t.Parallel()

	n := 0
	screen := counterPage(t, &n)
	if err := screen.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	if _, ok := screen.TakeDirty(); ok {
		t.Fatal("identical redraw reported dirty")
	}
}

func TestClickDiffFindsCountBox(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	n := 0
	screen := counterPage(t, &n)
	x, y := boxCenter(t, screen, "inc")
	if err := screen.Click(ctx, x, y); err != nil {
		t.Fatal(err)
	}

	rect, ok := screen.TakeDirty()
	if !ok {
		t.Fatal("no dirty rect")
	}

	insideBox(t, rect, boxFor(t, screen, "count"), 8)
}
