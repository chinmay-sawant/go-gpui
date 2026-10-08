// Package bitmap paints a display list into an image. blinkless returns the
// list and no page bitmap. This painter covers filled rectangles so a PNG
// can be encoded outside an Ebiten game. The window still replays the list.
package bitmap

import (
	"image"
	"image/color"
	"math"

	"github.com/chinmay-sawant/blinkless/layout"
)

// pxPerPt is the CSS pixel length of one layout point at zoom 1.
const pxPerPt = 96.0 / 72.0

// Picture paints the fill operations in display. A nil or empty canvas
// returns nil.
func Picture(display *layout.Display) image.Image {
	if display == nil || display.Width < 1 || display.Height < 1 {
		return nil
	}

	img := image.NewNRGBA(image.Rect(0, 0, display.Width, display.Height))
	clearWhite(img)
	faces := faceCache{}

	for _, index := range display.Order {
		if index < 0 || index >= len(display.Ops) {
			continue
		}

		op := &display.Ops[index]
		switch op.Kind {
		case layout.DisplayOpFillRect:
			fill(img, op)
		case layout.DisplayOpImage:
			paintImage(img, op)
		case layout.DisplayOpText, layout.DisplayOpBullet:
			paintText(img, op, faces)
		}
	}

	return img
}

func fill(img *image.NRGBA, op *layout.DisplayOp) {
	x, y := origin(op)
	left := int(math.Round(x * pxPerPt))
	top := int(math.Round(y * pxPerPt))
	right := int(math.Round((x + op.W) * pxPerPt))
	bottom := int(math.Round((y + op.H) * pxPerPt))
	bounds := img.Bounds()

	if left < bounds.Min.X {
		left = bounds.Min.X
	}

	if top < bounds.Min.Y {
		top = bounds.Min.Y
	}

	if right > bounds.Max.X {
		right = bounds.Max.X
	}

	if bottom > bounds.Max.Y {
		bottom = bounds.Max.Y
	}

	c := color.NRGBA{R: channel(op.R), G: channel(op.G), B: channel(op.B), A: 255}

	for y := top; y < bottom; y++ {
		for x := left; x < right; x++ {
			img.SetNRGBA(x, y, c)
		}
	}
}

func channel(value float64) uint8 {
	if value >= 1 {
		return 255
	}

	if value <= 0 {
		return 0
	}

	return uint8(math.Round(value * 255))
}
