package page

import (
	"bytes"
	"context"
	"image"
	"image/png"
	"strings"

	"github.com/chinmay-sawant/go-gpui/internal/render"
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

	img, boxes, err := render.Paint(ctx, body.String(), p.width, p.height)
	if err != nil {
		return err
	}

	p.img = img
	p.boxes = boxes
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
