package page

import "context"

// SavePDF is the package-level form of Page.SavePDF.
func SavePDF(ctx context.Context, p *Page, path string, opts PDFOptions) error {
	if p == nil {
		return ErrNilPage
	}

	return p.SavePDF(ctx, path, opts)
}

// pdfReady rejects a nil page and a cancelled context before a PDF call.
func pdfReady(ctx context.Context, p *Page) error {
	if p == nil {
		return ErrNilPage
	}

	return useContext(ctx)
}
