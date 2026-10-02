package replay

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/chinmay-sawant/gowkhtmltopdf/layout"
)

// drawImage draws one decoded image op into its layout rectangle. The payload
// is decoded once and kept in the image cache, so later frames reuse the same
// Ebiten image. A nil decode (empty payload or bad bytes) skips the op. The
// replay gate rejects non-identity transforms, so the placement is the op
// rectangle scaled from the decoded pixel bounds.
func drawImage(dst *ebiten.Image, op *layout.DisplayOp, dx, dy float64) {
	src := decodedImage(op)
	if src == nil {
		return
	}

	bounds := src.Bounds()

	options := &ebiten.DrawImageOptions{
		GeoM: imageGeom(op, dx, dy, bounds.Dx(), bounds.Dy()),
	}

	if alpha := op.Opacity(); alpha < 1 {
		options.ColorScale.ScaleAlpha(float32(alpha))
	}

	dst.DrawImage(src, options)
}
