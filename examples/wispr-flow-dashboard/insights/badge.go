package insights

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"

	xdraw "golang.org/x/image/draw"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/opentype"
)

const (
	badgeSize  = 70
	badgeScale = 4
)

// ringText repeats around the badge, one rune per step.
const ringText = "SHARE · SHARE · SHARE · "

// ShareBadgePNG draws the circular SHARE badge: a cream disc with the word
// repeated around a ring. The share arrow is an SVG painted on top.
func ShareBadgePNG() []byte {
	side := badgeSize * badgeScale
	img := image.NewRGBA(image.Rect(0, 0, side, side))
	cream := color.RGBA{R: 245, G: 244, B: 240, A: 255}

	drawDisc(img, float64(side)/2, float64(side)/2, float64(side)/2, cream)
	drawRing(img)

	small := image.NewRGBA(image.Rect(0, 0, badgeSize, badgeSize))
	xdraw.CatmullRom.Scale(small, small.Bounds(), img, img.Bounds(), draw.Src, nil)

	var buf bytes.Buffer
	_ = png.Encode(&buf, small)

	return buf.Bytes()
}

// drawDisc fills one circle.
func drawDisc(img *image.RGBA, cx, cy, r float64, c color.RGBA) {
	for y := range img.Bounds().Dy() {
		for x := range img.Bounds().Dx() {
			dx, dy := float64(x)-cx, float64(y)-cy
			if dx*dx+dy*dy <= r*r {
				img.SetRGBA(x, y, c)
			}
		}
	}
}

// drawRing places each ring rune upright against the rim.
func drawRing(img *image.RGBA) {
	ft, err := opentype.Parse(gobold.TTF)
	if err != nil {
		return
	}

	face, err := opentype.NewFace(ft, &opentype.FaceOptions{
		Size: 8 * badgeScale,
		DPI:  72,
	})
	if err != nil {
		return
	}
	defer func() { _ = face.Close() }()

	ink := color.RGBA{R: 10, G: 79, B: 69, A: 255}
	side := float64(img.Bounds().Dx())
	radius := side * 0.37
	step := 2 * math.Pi / float64(len(ringText))

	for i, r := range ringText {
		if r == ' ' {
			continue
		}

		angle := float64(i) * step
		x := side/2 + radius*math.Sin(angle)
		y := side/2 - radius*math.Cos(angle)

		rotateInto(img, drawGlyph(face, string(r), ink), x, y, angle)
	}
}
