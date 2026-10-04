package page_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

func TestHoverDirtiesBothBoxes(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	screen, err := page.New(page.Config{
		HTML: `<body style="margin:0">
<div id="a" style="width:80px;height:40px;background:#eee">a</div>
<div id="b" style="width:80px;height:40px;background:#ddd">b</div>
</body>`,
		Width:  320,
		Height: 200,
	})
	if err != nil {
		t.Fatal(err)
	}

	if err = screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}
	screen.TakeDirty()

	ax, ay := boxCenter(t, screen, "a")
	if err = screen.Hover(ctx, ax, ay); err != nil {
		t.Fatal(err)
	}

	if _, ok := screen.TakeDirty(); !ok {
		t.Fatal("hover into a not dirty")
	}

	bx, by := boxCenter(t, screen, "b")
	if err = screen.Hover(ctx, bx, by); err != nil {
		t.Fatal(err)
	}

	rect, ok := screen.TakeDirty()
	if !ok {
		t.Fatal("hover into b not dirty")
	}

	if !boxIn(rect, boxFor(t, screen, "a")) || !boxIn(rect, boxFor(t, screen, "b")) {
		t.Fatalf("rect %v misses a hovered box", rect)
	}
}
