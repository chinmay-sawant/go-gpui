package page_test

import (
	"context"
	"errors"
	"testing"

	"github.com/chinmay-sawant/ownframe/internal/page"
)

const reportHTML = `<html><head><title>Report</title></head><body>` +
	`<p>First paragraph.</p><p>Second paragraph.</p></body></html>`

// newReport builds the two-paragraph page the print tests use.
func newReport(t *testing.T) *page.Page {
	t.Helper()

	p, err := page.New(page.Config{HTML: reportHTML, Width: 640, Height: 480})
	if err != nil {
		t.Fatal(err)
	}

	return p
}

func TestPDFReportsNoWriter(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p := newReport(t)

	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	_, err := p.PDF(ctx, page.PDFOptions{})
	if !errors.Is(err, page.ErrNoPDF) {
		t.Fatalf("err = %v, want ErrNoPDF", err)
	}
}

func TestPDFWithThemeReportsNoWriter(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p, err := page.New(page.Config{
		HTML:   reportHTML,
		Theme:  `p{color:#ff0000}`,
		Width:  640,
		Height: 480,
	})
	if err != nil {
		t.Fatal(err)
	}

	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	if _, err := p.PDF(ctx, page.PDFOptions{}); !errors.Is(err, page.ErrNoPDF) {
		t.Fatalf("err = %v, want ErrNoPDF", err)
	}
}
