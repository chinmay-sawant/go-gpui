package ui

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestCSVImportPreviewAndCommit(t *testing.T) {
	b := newFake()
	app := newTestApp(t, b)
	settle(t, app)

	clickAction(t, app, "import")
	if !app.csv.open || app.csv.export {
		t.Fatalf("dialog = %+v", app.csv)
	}

	path := filepath.Join(t.TempDir(), "in.csv")
	if err := os.WriteFile(path, []byte("a,b\nc,d\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	app.loadCSV(path)
	flush(t, app)
	if !app.csv.ready || app.csv.prev.Total != 2 || app.csv.prev.Cols != 2 {
		t.Fatalf("preview = %+v", app.csv)
	}

	before := app.rev
	app.commitCSV(context.Background())
	flush(t, app)
	if app.csv.open {
		t.Fatal("dialog stayed open after commit")
	}

	if app.rev == before {
		t.Fatal("revision did not change")
	}
}

func TestCSVExportWritesFile(t *testing.T) {
	b := newFake()
	app := newTestApp(t, b)
	settle(t, app)

	clickAction(t, app, "export")
	if !app.csv.open || !app.csv.export {
		t.Fatalf("dialog = %+v", app.csv)
	}

	path := filepath.Join(t.TempDir(), "out.csv")
	app.csv.path = path
	app.writeCSV(context.Background())
	flush(t, app)

	if _, err := os.Stat(path); err != nil {
		t.Fatal(err)
	}

	if app.csv.open {
		t.Fatal("export dialog stayed open")
	}
}

func TestCSVCommitWithoutFileKeepsWorkbook(t *testing.T) {
	app := newTestApp(t, newFake())
	settle(t, app)

	clickAction(t, app, "import")
	before := app.rev
	app.commitCSV(context.Background())
	flush(t, app)

	if app.rev != before {
		t.Fatal("commit without data changed the revision")
	}

	if app.status == "" {
		t.Fatal("no status for the empty commit")
	}
}
