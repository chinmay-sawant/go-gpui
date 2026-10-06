package resize_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/resize/resize"
)

// TestMinimumSizeKeepsTheButtonInItsColumn shrinks the frame below the
// declared minimum. The size clamps up, and the hover button stays inside
// its column at the smallest frame instead of overflowing it.
func TestMinimumSizeKeepsTheButtonInItsColumn(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app, err := resize.New()
	if err != nil {
		t.Fatal(err)
	}

	app.SetSize(resize.MinWidth-100, resize.MinHeight-50)

	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	w, h := app.Page().Size()
	if w != resize.MinWidth || h != resize.MinHeight {
		t.Fatalf("size = %dx%d, want the %dx%d minimum", w, h, resize.MinWidth, resize.MinHeight)
	}

	btn := boxOf(t, app, "hover-btn")
	left := boxOf(t, app, "left")

	if right := btn.X + btn.W; right > left.X+left.W+0.5 {
		t.Fatalf("button right = %.0f, past the column right %.0f", right, left.X+left.W)
	}

	if right := btn.X + btn.W; right > float64(w)+0.5 {
		t.Fatalf("button right = %.0f, past the %d px frame", right, w)
	}
}
