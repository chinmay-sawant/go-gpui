package page_test

import (
	"bytes"
	"context"
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

func TestPDFStartsWithHeader(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	p := newReport(t)

	if err := p.Redraw(ctx); err != nil {
		t.Fatal(err)
	}

	data, err := p.PDF(ctx, page.PDFOptions{})
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.HasPrefix(data, []byte("%PDF-")) {
		t.Fatalf("PDF starts with %q", data[:min(8, len(data))])
	}
}

func TestPDFTwoParagraphsIsOnePage(t *testing.T) {
	t.Parallel()

	data, err := newReport(t).PDF(context.Background(), page.PDFOptions{})
	if err != nil {
		t.Fatal(err)
	}

	pages := bytes.Count(data, []byte("/Type /Page")) - bytes.Count(data, []byte("/Type /Pages"))
	if pages != 1 {
		t.Fatalf("pages = %d, want 1", pages)
	}
}

func TestPDFWithTheme(t *testing.T) {
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

	if _, err := p.PDF(ctx, page.PDFOptions{}); err != nil {
		t.Fatal(err)
	}
}
