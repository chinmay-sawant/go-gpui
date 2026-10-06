package ui

import (
	"context"
	"fmt"
	"testing"
)

// benchRows builds n running rows with distinct IDs.
func benchRows(n int) []Row {
	rows := make([]Row, 0, n)

	for i := 0; i < n; i++ {
		row := running(fmt.Sprintf("job-%02d", i), int64(i*1000), 1_000_000)
		row.Name = fmt.Sprintf("benchmark-file-%02d.bin", i)
		rows = append(rows, row)
	}

	return rows
}

// BenchmarkActiveTickPaint measures a tick that only edits retained
// operations: one progress update for a 20-row queue.
func BenchmarkActiveTickPaint(b *testing.B) {
	app, back := newTestApp(b)
	ctx := context.Background()
	rows := benchRows(20)

	back.send(Update{Kind: UpdateActive, Active: rows})
	if err := app.Tick(ctx); err != nil {
		b.Fatalf("Tick: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		row := rows[i%len(rows)]
		row.Done = int64(i % 100)
		back.send(Update{Kind: UpdateProgress, Row: &row})

		if err := app.Tick(ctx); err != nil {
			b.Fatalf("Tick: %v", err)
		}
	}
}

// BenchmarkActiveTickRedraw measures a tick that relayouts first, the cost
// a state change or resize pays.
func BenchmarkActiveTickRedraw(b *testing.B) {
	app, back := newTestApp(b)
	ctx := context.Background()
	rows := benchRows(20)

	back.send(Update{Kind: UpdateActive, Active: rows})
	if err := app.Tick(ctx); err != nil {
		b.Fatalf("Tick: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rows[i%len(rows)].Done = int64(i % 100)
		back.send(Update{Kind: UpdateActive, Active: rows})

		if err := app.Tick(ctx); err != nil {
			b.Fatalf("Tick: %v", err)
		}
	}
}

// BenchmarkRedrawFull measures one full relayout of the 20-row queue.
func BenchmarkRedrawFull(b *testing.B) {
	app, back := newTestApp(b)
	ctx := context.Background()

	back.send(Update{Kind: UpdateActive, Active: benchRows(20)})
	if err := app.Tick(ctx); err != nil {
		b.Fatalf("Tick: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := app.Redraw(ctx); err != nil {
			b.Fatalf("Redraw: %v", err)
		}
	}
}
