package player

import (
	"bytes"
	"context"
	"image"
	_ "image/jpeg" // registers JPEG for DecodeConfig
	_ "image/png"  // registers PNG for DecodeConfig
	"strings"

	"github.com/chinmay-sawant/go-gpui"
)

// coverResult is one fetched artwork body. A nil data field means the slot
// keeps its placeholder.
type coverResult struct {
	name string
	data []byte
}

// loadCovers fetches artwork concurrently and registers every body that
// decodes as a PNG or JPEG. Only this goroutine calls SetImage.
func (a *App) loadCovers(ctx context.Context, songs []itunesSong, tracks []Track) {
	results := make(chan coverResult, len(tracks))
	pending := 0

	for i, song := range songs {
		if i >= len(tracks) || i >= maxCovers {
			break
		}

		raw := upscale(song.ArtworkURL100)
		if raw == "" {
			continue
		}

		pending++
		go fetchCover(ctx, tracks[i].Cover, raw, results)
	}

	for ; pending > 0; pending-- {
		res := <-results
		if res.data != nil {
			a.page.SetImage(res.name, res.data)
		}
	}
}

// fetchCover sends one artwork request and forwards only valid image bytes.
func fetchCover(ctx context.Context, name, raw string, out chan<- coverResult) {
	res, err := gpui.Fetch(ctx, raw)
	if err != nil || res.Status != 200 || !validImage(res.Body) {
		out <- coverResult{name: name}

		return
	}

	out <- coverResult{name: name, data: res.Body}
}

// validImage reports whether data is a PNG or JPEG with real dimensions.
func validImage(data []byte) bool {
	cfg, format, err := image.DecodeConfig(bytes.NewReader(data))
	if err != nil || cfg.Width <= 0 || cfg.Height <= 0 {
		return false
	}

	return format == "png" || format == "jpeg"
}

// upscale replaces the final 100x100bb segment of an artwork URL with
// 300x300bb, the size the player covers draw at.
func upscale(raw string) string {
	const tail = "100x100bb."

	i := strings.LastIndex(raw, tail)
	if i < 0 || strings.LastIndex(raw, "/") > i {
		return raw
	}

	return raw[:i] + "300x300bb." + raw[i+len(tail):]
}
