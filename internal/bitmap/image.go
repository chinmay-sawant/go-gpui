package bitmap

import (
	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"math"

	"github.com/chinmay-sawant/blinkless/layout"
)

func clearWhite(img *image.NRGBA) {
	white := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	for i := 0; i < len(img.Pix); i += 4 {
		img.Pix[i] = white.R
		img.Pix[i+1] = white.G
		img.Pix[i+2] = white.B
		img.Pix[i+3] = white.A
	}
}

func paintImage(dst *image.NRGBA, op *layout.DisplayOp) {
	data, _, _ := op.ImageBytes()
	if len(data) == 0 {
		return
	}

	src, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		src, _, err = image.Decode(bytes.NewReader(data))
		if err != nil {
			return
		}
	}

	x, y := origin(op)
	left := int(math.Round(x * pxPerPt))
	top := int(math.Round(y * pxPerPt))
	right := int(math.Round((x + op.W) * pxPerPt))
	bottom := int(math.Round((y + op.H) * pxPerPt))
	dw, dh := right-left, bottom-top
	if dw < 1 || dh < 1 {
		return
	}

	sb := src.Bounds()
	sw, sh := sb.Dx(), sb.Dy()
	if sw < 1 || sh < 1 {
		return
	}

	clip := dst.Bounds()
	for y := top; y < bottom; y++ {
		if y < clip.Min.Y || y >= clip.Max.Y {
			continue
		}

		sy := sb.Min.Y + (y-top)*sh/dh
		for x := left; x < right; x++ {
			if x < clip.Min.X || x >= clip.Max.X {
				continue
			}

			sx := sb.Min.X + (x-left)*sw/dw
			dst.Set(x, y, src.At(sx, sy))
		}
	}
}
