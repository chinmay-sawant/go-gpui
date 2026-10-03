package ipc_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/go-gpui/examples/ipc/ipc"
)

// click finds the box with id and clicks its center.
func click(t *testing.T, ctx context.Context, app *ipc.App, id string) {
	t.Helper()

	var x, y float64
	found := false

	for _, b := range app.Boxes() {
		if b.ID != id || b.W <= 0 || b.H <= 0 {
			continue
		}

		x = b.X + b.W/2
		y = b.Y + b.H/2
		found = true
	}

	if !found {
		t.Fatalf("no box id=%q", id)
	}

	if err := app.Click(ctx, x, y); err != nil {
		t.Fatal(err)
	}
}
