// Package music fetches free music from the Openverse API and plays it
// through the OS audio device. It is example-side support: the audio player
// and Spotify player examples import it, and package gpui itself does not.
package music

import (
	"context"
	"errors"
)

// Clip is one piece of playable free music.
type Clip struct {
	Title   string
	Creator string
	License string // "CC BY-SA 3.0", or "generated" for the demo tune
	URL     string
	Data    []byte // encoded MP3 or WAV bytes
}

// Resolver turns a search query and a pick number into playable audio.
type Resolver interface {
	Resolve(ctx context.Context, query string, pick int) (Clip, error)
}

// Library tries a primary source and falls back to a generated demo tune, so
// a page always has something to play.
type Library struct {
	Primary Resolver
}

// Resolve returns the primary result, or a demo tone when it fails.
func (l *Library) Resolve(ctx context.Context, query string, pick int) (Clip, error) {
	if l.Primary != nil {
		if clip, err := l.Primary.Resolve(ctx, query, pick); err == nil {
			return clip, nil
		}
	}

	return DemoTune(query), nil
}

// ErrNoAudio means no Ebiten audio context is running.
var ErrNoAudio = errors.New("music: no audio context; run the page through gpui.Run")
