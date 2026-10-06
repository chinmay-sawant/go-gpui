package page_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/page"
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

	display := screen.Display()
	if display == nil {
		t.Fatal("display is nil")
	}

	if display.Height <= 400 {
		t.Fatalf("display height = %d, want the content past the window", display.Height)
	}
}
