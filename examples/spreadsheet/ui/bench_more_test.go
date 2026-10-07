package ui

import (
	"context"
	"testing"
)

// BenchmarkRedrawWindowReplace crosses an overscan edge every step.
func BenchmarkRedrawWindowReplace(b *testing.B) {
	app := benchApp(b)
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		app.page.SetScrollOffset(0, (i%40)*RowH*8)
		app.page.StepScrollWindow()
		if err := app.page.Redraw(ctx); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkRedrawSameSource relayouts the cached document with no change.
func BenchmarkRedrawSameSource(b *testing.B) {
	app := benchApp(b)
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := app.page.Redraw(ctx); err != nil {
			b.Fatal(err)
		}
	}
}
