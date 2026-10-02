package render_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/render"
)

// TestDisplayIsReachableFromGoGpui is the guard for the display-list export.
// The operations live behind the engine's internal/ rule, so the published
// module cannot hand them over. This test fails to build if that export goes
// away or if the replace directive in go.mod is dropped.
func TestDisplayIsReachableFromGoGpui(t *testing.T) {
	t.Parallel()

	display, err := render.DisplayList(
		context.Background(),
		`<h1 style="color:#036">Sign in</h1><p id="hi" data-action="go">Hello</p>`,
		320,
		200,
	)
	if err != nil {
		t.Fatal(err)
	}

	if len(display.Ops) == 0 {
		t.Fatal("no ops")
	}

	if len(display.Order) != len(display.Ops) {
		t.Fatalf("order covers %d of %d ops", len(display.Order), len(display.Ops))
	}

	if display.Width <= 0 || display.Height <= 0 {
		t.Fatalf("canvas %dx%d", display.Width, display.Height)
	}
}

// TestDisplayAgreesWithPaintOnCanvas keeps the two entries consistent. A
// caller switching from Paint to DisplayList must not see its window geometry
// change, because both start from one placement.
func TestDisplayAgreesWithPaintOnCanvas(t *testing.T) {
	t.Parallel()

	const source = `<h1>Heading</h1><p>Body copy long enough to wrap.</p>`

	img, _, err := render.Paint(context.Background(), source, 320, 200)
	if err != nil {
		t.Fatal(err)
	}

	display, err := render.DisplayList(context.Background(), source, 320, 200)
	if err != nil {
		t.Fatal(err)
	}

	bounds := img.Bounds()
	if bounds.Dx() != display.Width || bounds.Dy() != display.Height {
		t.Fatalf("paint %dx%d, display %dx%d",
			bounds.Dx(), bounds.Dy(), display.Width, display.Height)
	}
}
