package bitmap

import (
	"image"
	"image/color"
	"math"
	"unsafe"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"

	"github.com/chinmay-sawant/blinkless/layout"
)

type faceKey struct {
	ptr  uintptr
	size float64
}

type faceCache map[faceKey]font.Face

func (c faceCache) face(data []byte, size float64) font.Face {
	key := faceKey{ptr: uintptr(unsafe.Pointer(&data[0])), size: size}
	if face, ok := c[key]; ok {
		return face
	}

	parsed, err := opentype.Parse(data)
	if err != nil {
		c[key] = nil

		return nil
	}

	face, err := opentype.NewFace(parsed, &opentype.FaceOptions{
		Size: size,
		DPI:  96,
	})
	if err != nil {
		c[key] = nil

		return nil
	}

	c[key] = face

	return face
}

func paintText(dst *image.NRGBA, op *layout.DisplayOp, cache faceCache) {
	if op.Font == nil || op.Text == "" {
		return
	}

	data := op.Font.Bytes()
	if len(data) == 0 || op.Size <= 0 {
		return
	}

	face := cache.face(data, op.Size)
	if face == nil {
		return
	}

	content := layout.DisplayTransformText(op.Text, op.TextTransformValue())
	x, y := origin(op)
	draw := font.Drawer{
		Dst: dst,
		Src: image.NewUniform(color.NRGBA{
			R: channel(op.R), G: channel(op.G), B: channel(op.B), A: 255,
		}),
		Face: face,
		Dot: fixed.Point26_6{
			X: fixed.Int26_6(math.Round(x * pxPerPt * 64)),
			Y: fixed.Int26_6(math.Round(y * pxPerPt * 64)),
		},
	}
	draw.DrawString(content)

	if layout.DisplayFakeBold(op) {
		draw.Dot.X += fixed.I(1)
		draw.DrawString(content)
	}
}
