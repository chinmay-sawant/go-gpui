package page_test

import (
	"context"
	"image"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

func TestFallbackRedrawDirtiesFrame(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	screen, err := page.New(page.Config{
		HTML:   `<style>div { width: 60px; height: 40px; background: #123456; border-radius: 50% / 20%; }</style><div>{{.V}}</div>`,
		Width:  320,
		Height: 400,
	})
	if err != nil {
		t.Fatal(err)
	}

	screen.SetData(struct{ V string }{V: "a"})
	if err = screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}
	screen.TakeDirty()

	screen.SetData(struct{ V string }{V: "b"})
	if err = screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	rect, ok := screen.TakeDirty()
	if !ok || rect != image.Rect(0, 0, 320, 400) {
		t.Fatalf("rect %v ok %v", rect, ok)
	}
}
