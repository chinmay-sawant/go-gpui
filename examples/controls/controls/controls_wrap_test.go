package controls_test

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/controls/controls"
)

// TestLongDocGrowsField checks that a long file name wraps in the doc field
// and grows the box instead of overflowing on one line.
func TestLongDocGrowsField(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app := newApp(t, ctx)
	before := fieldHeight(t, app, "doc")

	click(t, ctx, app, "doc")
	if err := app.Type(ctx, "some/very/long/path/to/a/file/that/keeps/going/document.pdf"); err != nil {
		t.Fatal(err)
	}

	if got := fieldHeight(t, app, "doc"); got <= before {
		t.Fatalf("doc height = %v, want more than %v", got, before)
	}
}

func fieldHeight(t *testing.T, app *controls.App, id string) float64 {
	t.Helper()

	for _, b := range app.Boxes() {
		if b.ID == id {
			return b.H
		}
	}

	t.Fatalf("no box id=%q", id)

	return 0
}
