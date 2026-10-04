package page_test

import (
	"context"
	"image"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

func TestInvalidateCountBoxOnly(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	n := 0
	screen := counterPage(t, &n)
	screen.Invalidate("count")

	if err := screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	rect, ok := screen.TakeDirty()
	if !ok {
		t.Fatal("no dirty rect")
	}

	count := boxFor(t, screen, "count")
	insideBox(t, rect, count, 8)

	if !boxIn(rect, count) {
		t.Fatalf("rect %v does not cover #count", rect)
	}
}

func TestInvalidateCoversDescendant(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	screen, err := page.New(page.Config{
		HTML:   `<body style="margin:0"><div id="count"><div id="inner">1</div></div></body>`,
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
	screen.Invalidate("count")

	if err = screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	rect, ok := screen.TakeDirty()
	if !ok {
		t.Fatal("no dirty rect")
	}

	if !boxIn(rect, boxFor(t, screen, "inner")) {
		t.Fatalf("rect %v does not cover the nested span", rect)
	}
}

func TestInvalidateMissingIsFullFrame(t *testing.T) {
	t.Parallel()

	n := 0
	screen := counterPage(t, &n)
	want := image.Rect(0, 0, 320, 200)

	for _, id := range []string{"", "missing"} {
		screen.Invalidate(id)

		rect, ok := screen.TakeDirty()
		if !ok || rect != want {
			t.Fatalf("id %q: rect %v ok %v", id, rect, ok)
		}
	}
}
