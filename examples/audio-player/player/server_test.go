package player

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
)

// cannedServer serves one canned search and the artwork it names. Artwork is
// served only under a 300x300bb tail, so a broken upscale cannot pass.
func cannedServer(t *testing.T) (*httptest.Server, *atomic.Int32) {
	t.Helper()

	var artwork atomic.Int32
	var srv *httptest.Server

	srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/search":
			fmt.Fprintf(w, `{"resultCount":3,"results":[%s,%s,%s]}`,
				songJSON("Canned One", "Test Artist", srv.URL+"/art/100x100bb.png"),
				songJSON("Canned Two", "Test Artist Two", srv.URL+"/art/100x100bb.jpeg"),
				`{"wrapperType":"track","kind":"podcast","trackName":"Skip Me"}`)
		case strings.HasPrefix(r.URL.Path, "/art/") && strings.Contains(r.URL.Path, "300x300bb"):
			artwork.Add(1)

			if strings.HasSuffix(r.URL.Path, ".png") {
				writePNG(w)
			} else {
				writeJPEG(w)
			}
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(srv.Close)

	return srv, &artwork
}

func songJSON(title, artist, art string) string {
	return fmt.Sprintf(`{"wrapperType":"track","kind":"song","trackName":%q,`+
		`"artistName":%q,"collectionName":"Test Album","primaryGenreName":"Pop",`+
		`"trackTimeMillis":185000,"artworkUrl100":%q}`, title, artist, art)
}

func writePNG(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "image/png")

	img := image.NewNRGBA(image.Rect(0, 0, 2, 2))
	for i := 0; i < len(img.Pix); i += 4 {
		img.Pix[i] = 255
		img.Pix[i+3] = 255
	}

	_ = png.Encode(w, img)
}

func writeJPEG(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "image/jpeg")

	img := image.NewRGBA(image.Rect(0, 0, 2, 2))
	for y := 0; y < 2; y++ {
		for x := 0; x < 2; x++ {
			img.Set(x, y, color.RGBA{B: 255, A: 255})
		}
	}

	_ = jpeg.Encode(w, img, nil)
}
