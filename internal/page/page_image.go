package page

import (
	"bytes"
	"context"
	"image"
	"image/png"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"

	"github.com/chinmay-sawant/go-gpui/internal/render"
)

// Image returns the last rasterized picture, or nil when nothing has been
// drawn or the page is drawn from its display list.
func (p *Page) Image() image.Image {
	return p.img
}

// Display returns the last display list, or nil when the page is drawn as a
// rasterized picture.
func (p *Page) Display() *layout.Display {
	return p.display
}

// PNG encodes the page to PNG bytes. A display-list page rasterizes on
// demand, once, and the bytes are cached until the next Redraw. It returns
// nil when the page cannot be painted.
func (p *Page) PNG() []byte {
	if p.png != nil {
		return p.png
	}

	img := p.img
	if img == nil {
		if p.source == "" {
			return nil
		}

		var err error

		img, _, err = render.Paint(context.Background(), p.source, p.width, p.height)
		if err != nil {
			return nil
		}
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil
	}

	p.png = buf.Bytes()

	return p.png
}

// Boxes returns the last hit-test boxes, or nil when nothing has been drawn.
func (p *Page) Boxes() []Box {
	return p.boxes
}
