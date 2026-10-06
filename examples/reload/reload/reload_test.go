package reload_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/reload/reload"
)

const testHTML = `<button id="bump">Count: {{.Count}}</button><input id="note" type="text">`

func center(t *testing.T, app *reload.App, id string) (float64, float64) {
	t.Helper()

	for _, b := range app.Boxes() {
		if b.ID == id && b.W > 0 && b.H > 0 {
			return b.X + b.W/2, b.Y + b.H/2
		}
	}

	t.Fatalf("no box id=%q", id)

	return 0, 0
}

func TestCounterClick(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app, err := reload.NewHTML(testHTML, false)
	if err != nil {
		t.Fatal(err)
	}

	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	x, y := center(t, app, "bump")

	for i := 0; i < 2; i++ {
		if err := app.Click(ctx, x, y); err != nil {
			t.Fatal(err)
		}
	}

	if app.View().Count != 2 {
		t.Fatalf("count = %d, want 2", app.View().Count)
	}
}

func TestNewFileWatchesTheCopy(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	path := filepath.Join(t.TempDir(), "index.html")
	if err := os.WriteFile(path, []byte(testHTML), 0o600); err != nil {
		t.Fatal(err)
	}

	app, err := reload.NewFile(path, true)
	if err != nil {
		t.Fatal(err)
	}

	if err := app.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if app.HTML() != testHTML {
		t.Fatalf("HTML = %q", app.HTML())
	}

	next := testHTML + `<p id="extra">extra</p>`
	if err := os.WriteFile(path, []byte(next), 0o600); err != nil {
		t.Fatal(err)
	}

	changed, err := app.PollReload(ctx)
	if !changed || err != nil {
		t.Fatalf("poll = %v, %v", changed, err)
	}

	if app.HTML() != next {
		t.Fatalf("HTML = %q", app.HTML())
	}
}
