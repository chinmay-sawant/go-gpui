package player

import (
	"context"
	"time"

	"github.com/chinmay-sawant/go-gpui/examples/music"
)

// fakeResolver answers every request with the same clip.
type fakeResolver struct{ clip music.Clip }

// Resolve returns the canned clip.
func (f fakeResolver) Resolve(context.Context, string, int) (music.Clip, error) {
	return f.clip, nil
}

// fakeVoice is an in-memory music.Voice for tests.
type fakeVoice struct {
	playing bool
	pos     time.Duration
	dur     time.Duration
	volume  float64
}

// Play starts the voice.
func (v *fakeVoice) Play() { v.playing = true }

// Pause stops the voice.
func (v *fakeVoice) Pause() { v.playing = false }

// Playing reports the play flag.
func (v *fakeVoice) Playing() bool { return v.playing }

// Position is the fake clock.
func (v *fakeVoice) Position() time.Duration { return v.pos }

// Duration is the fake length.
func (v *fakeVoice) Duration() time.Duration { return v.dur }

// Seek moves the fake clock.
func (v *fakeVoice) Seek(at time.Duration) error {
	v.pos = at

	return nil
}

// SetVolume records the level.
func (v *fakeVoice) SetVolume(volume float64) { v.volume = volume }

// Close is a no-op.
func (v *fakeVoice) Close() error { return nil }

// testClip is the clip the fake resolver returns.
var testClip = music.Clip{
	Title:   "Test clip",
	Creator: "Tester",
	License: "CC0 1.0",
	Data:    []byte("x"),
}
