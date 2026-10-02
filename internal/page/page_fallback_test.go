package page_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
)

func TestRoundedBorderReplays(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	screen, err := page.New(page.Config{
		HTML:   `<style>div { width: 60px; height: 40px; border: 2px solid #123456; border-radius: 6px; }</style><div></div>`,
		Width:  320,
		Height: 400,
	})
	if err != nil {
		t.Fatal(err)
	}

	if err := screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if screen.Display() == nil {
		t.Fatal("rounded border did not replay")
	}
}

func TestEllipticalFillFallsBackToBitmap(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	screen, err := page.New(page.Config{
		HTML:   `<style>div { width: 60px; height: 40px; background: #123456; border-radius: 50% / 20%; }</style><div></div>`,
		Width:  320,
		Height: 400,
	})
	if err != nil {
		t.Fatal(err)
	}

	if err := screen.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if screen.Display() != nil {
		t.Fatal("elliptical border was replayed")
	}

	if screen.Image() == nil {
		t.Fatal("no fallback bitmap")
	}
}
