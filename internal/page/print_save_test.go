package page_test

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/page"
)

func TestPDFRejectsUnknownProfile(t *testing.T) {
	t.Parallel()

	_, err := newReport(t).PDF(context.Background(), page.PDFOptions{Profile: "not-a-profile"})
	if !errors.Is(err, page.ErrNoPDF) {
		t.Fatalf("err = %v, want ErrNoPDF", err)
	}
}

func TestSavePDFWritesNoFile(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "report.pdf")

	err := newReport(t).SavePDF(context.Background(), path, page.PDFOptions{PageSize: "Letter"})
	if !errors.Is(err, page.ErrNoPDF) {
		t.Fatalf("err = %v, want ErrNoPDF", err)
	}

	if _, statErr := os.Stat(path); !os.IsNotExist(statErr) {
		t.Fatalf("stat = %v, want a missing file", statErr)
	}
}

func TestSavePDFNilPage(t *testing.T) {
	t.Parallel()

	err := page.SavePDF(context.Background(), nil, "unused.pdf", page.PDFOptions{})
	if !errors.Is(err, page.ErrNilPage) {
		t.Fatalf("err = %v, want ErrNilPage", err)
	}
}
