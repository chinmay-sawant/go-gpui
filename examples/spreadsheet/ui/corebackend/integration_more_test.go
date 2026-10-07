package corebackend

import (
	"context"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/ui"
)

func TestUICtrlEndJumpsToUsedRange(t *testing.T) {
	b := newBackend(t, t.TempDir())
	app, err := ui.New(ui.Options{Backend: b, Width: 1000, Height: 700, DataDir: t.TempDir()})
	if err != nil {
		t.Fatal(err)
	}

	defer app.Close()

	ctx := context.Background()
	waitFor(t, app, func() bool { return hasCell(app, "row 1") })

	if err := app.Page().KeyDown(ctx, "ctrl+end"); err != nil {
		t.Fatal(err)
	}

	waitFor(t, app, func() bool { return app.View().Ref == "T200" })
}

func TestUITypeCommitUndoOverStore(t *testing.T) {
	b := newBackend(t, t.TempDir())
	app, err := ui.New(ui.Options{Backend: b, Width: 1000, Height: 700, DataDir: t.TempDir()})
	if err != nil {
		t.Fatal(err)
	}

	defer app.Close()

	ctx := context.Background()
	page := app.Page()
	waitFor(t, app, func() bool { return hasCell(app, "row 1") })

	// Replace A1 and commit with Enter.
	if err := page.Type(ctx, "edited"); err != nil {
		t.Fatal(err)
	}

	if err := page.Submit(ctx); err != nil {
		t.Fatal(err)
	}

	id := sheetID(t, b, "Numbers")
	waitFor(t, app, func() bool {
		cells, err := b.Range(id, ui.Area{R0: 0, C0: 0, R1: 0, C1: 0})

		return err == nil && cells[0].Raw == "edited"
	})

	// Undo through the Ctrl chord handler.
	if err := page.Undo(ctx); err != nil {
		t.Fatal(err)
	}

	waitFor(t, app, func() bool {
		cells, err := b.Range(id, ui.Area{R0: 0, C0: 0, R1: 0, C1: 0})

		return err == nil && cells[0].Raw == "row 1"
	})
}
