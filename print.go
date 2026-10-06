package ownframe

import (
	"context"

	"github.com/chinmay-sawant/ownframe/internal/page"
)

// PDFOptions are the page settings for a PDF export. A zero value uses the
// engine defaults: A4, 10 mm margins, and no PDF profile.
type PDFOptions = page.PDFOptions

// ErrNoPrinter means the system has no print path. Page.Print wraps it and
// names SavePDF as the fallback.
var ErrNoPrinter = page.ErrNoPrinter

// SavePDF writes the page's last template output to path as a PDF.
func SavePDF(ctx context.Context, p *Page, path string, opts PDFOptions) error {
	return page.SavePDF(ctx, p, path, opts)
}
