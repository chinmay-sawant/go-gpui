package telegram_test

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"testing"
)

func solidPNG(t *testing.T, w, h int, c color.RGBA) []byte {
	t.Helper()

	img := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(img, img.Bounds(), &image.Uniform{C: c}, image.Point{}, draw.Src)

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		t.Fatal(err)
	}

	return buf.Bytes()
}

func toRGBA(t *testing.T, data []byte) *image.RGBA {
	t.Helper()

	src, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	b := src.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), src, b.Min, draw.Src)

	return dst
}

func crop(img image.Image, r image.Rectangle) *image.RGBA {
	dst := image.NewRGBA(image.Rect(0, 0, r.Dx(), r.Dy()))
	draw.Draw(dst, dst.Bounds(), img, r.Min, draw.Src)

	return dst
}

func samePix(a, b *image.RGBA) bool {
	return a.Bounds() == b.Bounds() && bytes.Equal(a.Pix, b.Pix)
}

// nearPix allows the rasterizer's anti-aliasing noise: at most maxN pixels
// differ, none by more than maxd on any channel.
func nearPix(a, b *image.RGBA, maxd, maxN int) bool {
	if a.Bounds() != b.Bounds() {
		return false
	}

	n, max := 0, 0

	for y := 0; y < a.Bounds().Dy(); y++ {
		for x := 0; x < a.Bounds().Dx(); x++ {
			ca, cb := a.RGBAAt(x, y), b.RGBAAt(x, y)
			if ca == cb {
				continue
			}

			n++

			for i, da := range []uint8{ca.R, ca.G, ca.B, ca.A} {
				db := []uint8{cb.R, cb.G, cb.B, cb.A}[i]
				d := int(da) - int(db)
				if d < 0 {
					d = -d
				}
				if d > max {
					max = d
				}
			}
		}
	}

	return n <= maxN && max <= maxd
}
