package print_test

import (
	"bytes"
	"context"
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

	data, err := app.Page().PDF(ctx, ownframe.PDFOptions{})
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.HasPrefix(data, []byte("%PDF-")) {
		t.Fatalf("PDF starts with %q", data[:8])
	}
}

func TestSaveButtonWritesPDF(t *testing.T) {
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

	data, err := os.ReadFile(app.SavePath)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.HasPrefix(data, []byte("%PDF-")) {
		t.Fatalf("saved file starts with %q", data[:8])
	}

	if got := app.View().Status; !strings.Contains(got, "Saved to") {
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
