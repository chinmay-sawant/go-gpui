package devtools

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
)

// shotPNG builds the small image the template shows through SetImage. The
// example ships no binary asset, so the bytes come from an encoder.
func shotPNG() []byte {
	img := image.NewRGBA(image.Rect(0, 0, 4, 4))

	for y := range 4 {
		for x := range 4 {
			img.Set(x, y, color.RGBA{
				R: uint8(0x30 * x),
				G: uint8(0x40 * y),
				B: 0xdb,
				A: 0xff,
			})
		}
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil
	}

	return buf.Bytes()
}
