package print_test

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/chinmay-sawant/ownframe"
	print "github.com/chinmay-sawant/ownframe/examples/print/print"
)

func TestReportPDF(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app, err := print.New(false)
	if err != nil {
		t.Fatal(err)
	}

	if err := app.Page().Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	_, err = app.Page().PDF(ctx, ownframe.PDFOptions{})
	if !errors.Is(err, ownframe.ErrNoPDF) {
		t.Fatalf("err = %v, want ErrNoPDF", err)
	}
}

func TestSaveButtonReportsNoWriter(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	app, err := print.New(false)
	if err != nil {
		t.Fatal(err)
	}

	app.SavePath = filepath.Join(t.TempDir(), "report.pdf")

	if err := app.Page().Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	click(t, ctx, app, "save")

	if _, statErr := os.Stat(app.SavePath); !os.IsNotExist(statErr) {
		t.Fatalf("stat = %v, want a missing file", statErr)
	}

	if got := app.View().Status; !strings.Contains(got, "Save failed") {
		t.Fatalf("status = %q", got)
	}
}

func click(t *testing.T, ctx context.Context, app *print.App, id string) {
	t.Helper()

	var x, y float64
	found := false

	for _, b := range app.Page().Boxes() {
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

	if err := app.Page().Click(ctx, x, y); err != nil {
		t.Fatal(err)
	}
}
