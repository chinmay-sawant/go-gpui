package ui

import (
	"context"
	"testing"
)

// benchApp builds a redrawn app with a fake source.
func benchApp(b *testing.B) *App {
	b.Helper()

	app, err := New(context.Background(), Config{
		Source: &fakeSource{},
		noPump: true,
		Width:  1280,
		Height: 720,
	})
	if err != nil {
		b.Fatal(err)
	}

	b.Cleanup(func() { _ = app.Close() })

	if err := app.page.Redraw(context.Background()); err != nil {
		b.Fatal(err)
	}

	return app
}

// BenchmarkTickPaint measures one active tick: drain, rebind check, and the
// bar and text paint on the retained display list.
func BenchmarkTickPaint(b *testing.B) {
	app := benchApp(b)
	ctx := context.Background()

	app.mail.putSummary(app.currentGen(), sample("cpu", 0.4))

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := app.Tick(ctx); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkRedraw measures a full repaint: parse is cached, cascade and
// layout run again. This is the cost a click or resize pays.
func BenchmarkRedraw(b *testing.B) {
	app := benchApp(b)
	ctx := context.Background()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := app.page.Redraw(ctx); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkRedrawProcesses measures a full repaint of the process screen
// with a 10,000 row snapshot, of which 50 rows render.
func BenchmarkRedrawProcesses(b *testing.B) {
	app := benchApp(b)
	ctx := context.Background()

	app.state.table.offer(snapshotOf(manyProcs(10000)...))
	app.state.table.refresh()
	app.state.nav = "processes"
	app.sync()

	if err := app.page.Redraw(ctx); err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := app.page.Redraw(ctx); err != nil {
			b.Fatal(err)
		}
	}
}
