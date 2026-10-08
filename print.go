package ownframe

import (
	"context"

	"github.com/chinmay-sawant/ownframe/internal/page"
)

// PDFOptions names a page size, a margin, and a PDF profile. blinkless does
// not write PDF, so the PDF methods return ErrNoPDF.
type PDFOptions = page.PDFOptions

// ErrNoPrinter means the system has no print path. Page.Print wraps it and
// names SavePDF as the fallback.
var ErrNoPrinter = page.ErrNoPrinter

// ErrNoPDF means the layout engine does not write a PDF.
var ErrNoPDF = page.ErrNoPDF

// SavePDF writes the page's last template output to path as a PDF.
func SavePDF(ctx context.Context, p *Page, path string, opts PDFOptions) error {
	return page.SavePDF(ctx, p, path, opts)
}
