package perfbench_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/examples/perf-benchmarks/benchutil"
)

// BenchmarkNormalRedraw measures a warm Redraw of the small desktop form:
// the data never changes, so this is the relayout plus layout of a
// mostly-idle page.
func BenchmarkNormalRedraw(b *testing.B) {
	ctx := context.Background()

	app, err := benchutil.NewNormal()
	if err != nil {
		b.Fatal(err)
	}

	if err := app.Redraw(ctx); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for b.Loop() {
		if err := app.Redraw(ctx); err != nil {
			b.Fatal(err)
		}
	}
}
