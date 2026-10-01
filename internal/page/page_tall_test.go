package page_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

func TestTallPagePaintsPastTheWindow(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	screen, err := page.New(page.Config{
		HTML:   `<style>body { margin: 0; }</style><div style="height: 1200px">tall</div>`,
		Width:  320,
		Height: 400,
	})
	if err != nil {
		t.Fatal(err)
	}

	if err := screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	height := screen.Image().Bounds().Dy()
	if height <= 400 {
		t.Fatalf("image height = %d, want the content past the window", height)
	}
}
