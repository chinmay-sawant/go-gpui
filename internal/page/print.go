package page

import (
	"context"
	"io"
	"os"

	"github.com/chinmay-sawant/ownframe/internal/print"
)

// PDFOptions are the page settings for a PDF export. A zero value uses the
// engine defaults: A4, 10 mm margins, and no PDF profile.
type PDFOptions struct {
	// PageSize is an engine page name such as "A4" or "Letter". Empty is A4.
	PageSize string
	// Margin is one width in millimetres applied to all four sides.
	// Zero is the engine default, 10 mm.
	Margin float64
	// Profile is a PDF conformance profile such as "PDF/A-4" or "PDF/UA-1".
	// Empty writes no profile. A profile the engine does not know is an
	// error.
	Profile string
}

// ErrNoPrinter means the system has no print path. Print wraps it and names
// SavePDF as the fallback.
var ErrNoPrinter = print.ErrNoPrinter

// PDF renders the last template output as a PDF and returns the bytes.
func (p *Page) PDF(ctx context.Context, opts PDFOptions) ([]byte, error) {
	doc, err := p.pdfDocument(ctx, opts)
	if err != nil {
		return nil, err
	}

	return doc.PDF(ctx)
}

// WritePDF renders the last template output as a PDF and writes it to w.
func (p *Page) WritePDF(ctx context.Context, w io.Writer, opts PDFOptions) error {
	doc, err := p.pdfDocument(ctx, opts)
	if err != nil {
		return err
	}

	return doc.WritePDF(ctx, w)
}

// SavePDF writes the last template output to path.
func (p *Page) SavePDF(ctx context.Context, path string, opts PDFOptions) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	err = p.WritePDF(ctx, file, opts)
	closeErr := file.Close()

	if err != nil {
		os.Remove(path)

		return err
	}

	return closeErr
}
