package perfbench_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/examples/perf-benchmarks/benchutil"
)

// BenchmarkFlappyTick measures one frame tick of the flappy demo: physics,
// collision, and moving the retained display-list ops, with no Redraw. The
// first Redraw binds the ops; the loop then times steady-state frames.
func BenchmarkFlappyTick(b *testing.B) {
	ctx := context.Background()

	app, err := benchutil.NewFlap()
	if err != nil {
		b.Fatal(err)
	}

	if err := app.Redraw(ctx); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for b.Loop() {
		if err := app.Tick(ctx); err != nil {
			b.Fatal(err)
		}
	}
}
