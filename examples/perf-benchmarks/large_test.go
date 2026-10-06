package perfbench_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/examples/perf-benchmarks/benchutil"
)

// BenchmarkLargeRedraw measures a warm Redraw of the thousand-row list while
// alternating the width by one pixel, so every iteration also pays the
// resize relayout a frequent-resize page must survive.
func BenchmarkLargeRedraw(b *testing.B) {
	ctx := context.Background()

	app, err := benchutil.NewLarge(1000)
	if err != nil {
		b.Fatal(err)
	}

	if err := app.Redraw(ctx); err != nil {
		b.Fatal(err)
	}

	widths := [2]int{800, 801}

	b.ResetTimer()

	for i := 0; b.Loop(); i++ {
		app.SetSize(widths[i%2], 600)

		if err := app.Redraw(ctx); err != nil {
			b.Fatal(err)
		}
	}
}
