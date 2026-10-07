package ui

import (
	"context"
	"testing"
)

func TestRevisionChangeInvalidatesTiles(t *testing.T) {
	app := newTestApp(t, newFake())
	settle(t, app)

	r := result{kind: jobApply, rev: app.rev + 1, edits: []Edit{{Row: 0, Col: 0, Raw: "x"}}}
	if !app.applyResult(r) {
		t.Fatal("edit ack did not redraw")
	}

	if app.rev != r.rev {
		t.Fatalf("rev = %d", app.rev)
	}
}

func TestTickBudget(t *testing.T) {
	app := newTestApp(t, newFake())
	settle(t, app)

	for range drainBudget + 10 {
		app.work.out <- result{kind: jobPref}
	}

	if err := app.tick(context.Background()); err != nil {
		t.Fatal(err)
	}

	left := len(app.work.out)
	if left != 10 {
		t.Fatalf("results left = %d, want 10", left)
	}
}
