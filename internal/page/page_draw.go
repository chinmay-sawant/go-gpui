package page

import (
	"bytes"
	"context"
	"image"
	"image/png"
	"strings"

	"github.com/chinmay-sawant/gowkhtmltopdf/css"
	"github.com/chinmay-sawant/gowkhtmltopdf/html"
	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// Redraw fills the template and renders the current size.
func (p *Page) Redraw(ctx context.Context) error {
	if err := useContext(ctx); err != nil {
		return err
	}

	var body strings.Builder
	if err := p.tpl.Execute(&body, p.data); err != nil {
		return err
	}

	doc, err := html.Parse([]byte(body.String()))
	if err != nil {
		return err
	}

	styled, err := css.Apply(ctx, doc, css.Options{
		WidthPx:  p.width,
		HeightPx: p.height,
		Media:    "screen",
		Extra:    nil,
	})
	if err != nil {
		return err
	}

	placed, err := layout.Lay(ctx, styled)
	if err != nil {
		return err
	}

	p.img = placed.Image()
	p.boxes = placed.Boxes()
	p.generation++

	return nil
}

// Image returns the last picture, or nil when nothing has been drawn.
func (p *Page) Image() image.Image {
	return p.img
}

// PNG encodes Image to PNG bytes.
// It returns nil when Image is nil.
func (p *Page) PNG() []byte {
	img := p.Image()
	if img == nil {
		return nil
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil
	}

	return buf.Bytes()
}

// Boxes returns the last hit-test boxes, or nil when nothing has been drawn.
func (p *Page) Boxes() []Box {
	return p.boxes
}
