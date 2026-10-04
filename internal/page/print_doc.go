package page

import (
	"context"
	"strings"

	"github.com/chinmay-sawant/gowkhtmltopdf"
)

// SavePDF is the package-level form of Page.SavePDF.
func SavePDF(ctx context.Context, p *Page, path string, opts PDFOptions) error {
	if p == nil {
		return ErrNilPage
	}

	return p.SavePDF(ctx, path, opts)
}

// pdfDocument wraps the print source in an engine document with opts applied.
func (p *Page) pdfDocument(ctx context.Context, opts PDFOptions) (*gowkhtmltopdf.Document, error) {
	source, err := p.pdfSource(ctx)
	if err != nil {
		return nil, err
	}

	doc := gowkhtmltopdf.NewDocument(gowkhtmltopdf.Page{
		Source: gowkhtmltopdf.HTML([]byte(source)),
	})
	doc.PageSize = opts.PageSize
	doc.PDFProfile = opts.Profile

	if opts.Margin > 0 {
		doc.Margin = gowkhtmltopdf.Margin{
			Top: opts.Margin, Right: opts.Margin, Bottom: opts.Margin, Left: opts.Margin,
		}
	}

	return doc, nil
}

// pdfSource returns the source the PDF renders: the last template output, or
// the current template when nothing has drawn yet, with the theme in a style
// element.
func (p *Page) pdfSource(ctx context.Context) (string, error) {
	if err := useContext(ctx); err != nil {
		return "", err
	}

	source := p.source
	if source == "" {
		var body strings.Builder
		if err := p.tpl.Execute(&body, p.data); err != nil {
			return "", err
		}

		source = p.syncForm(body.String())
	}

	return injectTheme(source, p.themeSrc), nil
}

// injectTheme puts theme in a style element in the head, before </head> when
// the source has one, and in front of the source otherwise.
func injectTheme(source, theme string) string {
	if strings.TrimSpace(theme) == "" {
		return source
	}

	style := "<style>" + theme + "</style>"
	lower := strings.ToLower(source)

	if at := strings.Index(lower, "</head>"); at >= 0 {
		return source[:at] + style + source[at:]
	}

	if at := strings.Index(lower, "<head>"); at >= 0 {
		at += len("<head>")

		return source[:at] + style + source[at:]
	}

	return style + source
}
