package assets

import (
	"bytes"
	"crypto/sha256"
	"image/png"
	"io/fs"
	"path"
	"strings"
	"testing"
)

func TestCatCollectionHasDistinctTransparentExpressions(t *testing.T) {
	files, err := fs.Glob(Cats, "cat_images/*.png")
	if err != nil {
		t.Fatal(err)
	}
	if len(files) < 30 {
		t.Fatalf("got %d expressions, want at least 30", len(files))
	}
	emotions := map[string]bool{}
	hashes := map[[32]byte]bool{}
	for _, name := range files {
		t.Run(path.Base(name), func(t *testing.T) {
			data, err := Cats.ReadFile(name)
			if err != nil {
				t.Fatal(err)
			}
			hash := sha256.Sum256(data)
			if hashes[hash] {
				t.Fatal("duplicate image bytes")
			}
			hashes[hash] = true
			img, err := png.Decode(bytes.NewReader(data))
			if err != nil {
				t.Fatal(err)
			}
			b := img.Bounds()
			if b.Dx() < 256 || b.Dy() < 256 {
				t.Fatal("source is too small for the desktop mascot")
			}
			for _, p := range [][2]int{{0, 0}, {b.Max.X - 1, 0}, {0, b.Max.Y - 1}, {b.Max.X - 1, b.Max.Y - 1}} {
				if _, _, _, alpha := img.At(p[0], p[1]).RGBA(); alpha != 0 {
					t.Fatal("PNG corner is not transparent")
				}
			}
			if _, _, _, alpha := img.At(b.Dx()/2, b.Dy()/2).RGBA(); alpha == 0 {
				t.Fatal("PNG has no visible cat at its center")
			}
		})
		emotions[strings.Split(path.Base(name), "-")[1]] = true
	}
	if len(emotions) < 10 || len(emotions) > 20 {
		t.Fatalf("got %d emotion groups, want 10 through 20", len(emotions))
	}
}
