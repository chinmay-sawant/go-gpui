package player

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"strings"
)

// tinyPNG and tinyJPEG are real image bytes for the artwork endpoints.
func tinyPNG() []byte {
	var buf bytes.Buffer
	img := image.NewRGBA(image.Rect(0, 0, 4, 4))
	_ = png.Encode(&buf, img)

	return buf.Bytes()
}

func tinyJPEG() []byte {
	var buf bytes.Buffer
	img := image.NewRGBA(image.Rect(0, 0, 4, 4))
	_ = jpeg.Encode(&buf, img, nil)

	return buf.Bytes()
}

// songsBody is a canned iTunes song response with one podcast decoy.
func songsBody(base string) string {
	var b strings.Builder

	b.WriteString(`{"results":[`)

	for i := 0; i < 6; i++ {
		if i > 0 {
			b.WriteString(",")
		}

		fmt.Fprintf(&b, `{"kind":"song","trackName":"Song %d","artistName":"Artist %d","collectionName":"Album %d","releaseDate":"2021-05-0%dT07:00:00Z","trackTimeMillis":%d000,"artworkUrl100":"%s/art/song-%d/100x100bb.png","primaryGenreName":"Pop"}`, i, i, i, i+1, 200+i, base, i)
	}

	fmt.Fprintf(&b, `,{"kind":"podcast","trackName":"Not A Song","artistName":"Talker","collectionName":"Podcast","releaseDate":"2020-01-01T00:00:00Z","trackTimeMillis":60000,"artworkUrl100":"%s/art/song-9/100x100bb.png"}]}`, base)

	return b.String()
}

// albumsBody is a canned iTunes album response.
func albumsBody(base string) string {
	var b strings.Builder

	b.WriteString(`{"results":[`)

	for i := 0; i < 12; i++ {
		if i > 0 {
			b.WriteString(",")
		}

		fmt.Fprintf(&b, `{"collectionName":"Record %d","artistName":"Band %d","artworkUrl100":"%s/art/album-%d/100x100bb.jpg"}`, i, i, base, i)
	}

	b.WriteString(`]}`)

	return b.String()
}
