package page_test

import (
	"context"
	"image"
	"testing"
)

func TestThemeDirtiesFullFrame(t *testing.T) {
	t.Parallel()

	n := 0
	screen := counterPage(t, &n)
	if err := screen.SetTheme("body { background: #000; }"); err != nil {
		t.Fatal(err)
	}

	if err := screen.Redraw(context.Background()); err != nil {
		t.Fatal(err)
	}

	rect, ok := screen.TakeDirty()
	if !ok || rect != image.Rect(0, 0, 320, 200) {
		t.Fatalf("rect %v ok %v", rect, ok)
	}
}
