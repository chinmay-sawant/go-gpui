package player

import (
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/frame"
)

// eqHeights reads every eq bar height in the head and the active row.
func eqHeights(t *testing.T, app *App) []float64 {
	t.Helper()

	d := app.Page().Display()
	if d == nil {
		t.Fatal("no display list")
	}

	var out []float64
	for _, id := range []string{"eq", "eq-row"} {
		box, ok := boxByID(app.Boxes(), id)
		if !ok {
			continue
		}

		for _, bar := range frame.Fills(d, box, accent) {
			out = append(out, bar.H)
		}
	}

	return out
}

// sameHeights reports whether two equal-length slices match exactly.
func sameHeights(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
