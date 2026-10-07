package corebackend

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/chinmay-sawant/ownframe/examples/spreadsheet/ui"
	pagepkg "github.com/chinmay-sawant/ownframe/internal/page"
)

func TestUICSVImportThroughPicker(t *testing.T) {
	b := newBackend(t, t.TempDir())
	app, err := ui.New(ui.Options{Backend: b, Width: 1000, Height: 700, DataDir: t.TempDir()})
	if err != nil {
		t.Fatal(err)
	}

	defer app.Close()

	page := app.Page()
	waitFor(t, app, func() bool { return hasCell(app, "row 1") })

	path := filepath.Join(t.TempDir(), "in.csv")
	if err := os.WriteFile(path, []byte("alpha,beta\n1,2\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	pagepkg.InstallPicker(page, func(context.Context, string) (string, bool) {
		return path, true
	})

	// Open the dialog from the toolbar, then click the file field.
	clickToolbar(t, app, "import")
	clickControl(t, app, "csvpath")
	waitFor(t, app, func() bool { return app.View().Dialog != nil && app.View().Dialog.Note != "" })

	// The preview keeps the workbook untouched until Commit.
	id := sheetID(t, b, "Numbers")
	if cells, _ := b.Range(id, ui.Area{R0: 0, C0: 0, R1: 0, C1: 0}); cells[0].Raw != "row 1" {
		t.Fatalf("preview changed the workbook: %+v", cells[0])
	}

	clickDialog(t, app, "csv-commit")
	waitFor(t, app, func() bool {
		cells, err := b.Range(id, ui.Area{R0: 0, C0: 0, R1: 1, C1: 1})

		return err == nil && cells[0].Raw == "alpha" && cells[3].Raw == "2"
	})
}
