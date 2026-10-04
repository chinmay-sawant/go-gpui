package page_test

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/chinmay-sawant/go-gpui/internal/page"
	"github.com/chinmay-sawant/gowkhtmltopdf"
)

func TestPDFRejectsUnknownProfile(t *testing.T) {
	t.Parallel()

	_, err := newReport(t).PDF(context.Background(), page.PDFOptions{Profile: "not-a-profile"})
	if err == nil {
		t.Fatal("unknown profile did not error")
	}

	if !errors.Is(err, gowkhtmltopdf.ErrInvalidPDFProfile) {
		t.Fatalf("err = %v, want ErrInvalidPDFProfile", err)
	}
}

func TestPDFAcceptsKnownProfile(t *testing.T) {
	t.Parallel()

	data, err := newReport(t).PDF(context.Background(), page.PDFOptions{Profile: "PDF/A-4"})
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.HasPrefix(data, []byte("%PDF-")) {
		t.Fatalf("profiled PDF starts with %q", data[:8])
	}
}

func TestSavePDFWritesFile(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "report.pdf")

	if err := newReport(t).SavePDF(context.Background(), path, page.PDFOptions{PageSize: "Letter"}); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.HasPrefix(data, []byte("%PDF-")) {
		t.Fatalf("file starts with %q", data[:8])
	}
}

func TestSavePDFNilPage(t *testing.T) {
	t.Parallel()

	err := page.SavePDF(context.Background(), nil, "unused.pdf", page.PDFOptions{})
	if !errors.Is(err, page.ErrNilPage) {
		t.Fatalf("err = %v, want ErrNilPage", err)
	}
}
