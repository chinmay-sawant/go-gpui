package insights

import (
	"image"
	"image/color"
	"math"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// drawGlyph renders one rune centered in a scratch image.
func drawGlyph(face font.Face, s string, ink color.RGBA) *image.RGBA {
	const box = 48

	img := image.NewRGBA(image.Rect(0, 0, box, box))
	d := font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(ink),
		Face: face,
	}
	width := d.MeasureString(s)
	d.Dot = fixed.Point26_6{X: (fixed.I(box) - width) / 2, Y: fixed.I(box/2 + 11)}
	d.DrawString(s)

	return img
}

// rotateInto places one rotated glyph into the badge image.
func rotateInto(dst, src *image.RGBA, cx, cy, angle float64) {
	sin, cos := math.Sincos(angle)
	w, h := src.Bounds().Dx(), src.Bounds().Dy()

	for y := range h {
		for x := range w {
			c := src.RGBAAt(x, y)
			if c.A == 0 {
				continue
			}

			sx := float64(x) - float64(w)/2
			sy := float64(y) - float64(h)/2
			dx := int(cx + sx*cos - sy*sin + 0.5)
			dy := int(cy + sx*sin + sy*cos + 0.5)

			if image.Pt(dx, dy).In(dst.Bounds()) {
				dst.SetRGBA(dx, dy, c)
			}
		}
	}
}
