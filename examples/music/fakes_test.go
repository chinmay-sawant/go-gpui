package music

import (
	"context"
	"time"
)

// fakeVoice records the calls an engine makes.
type fakeVoice struct {
	playing bool
	pos     time.Duration
	dur     time.Duration
	volume  float64
	closed  bool
}

func (f *fakeVoice) Play()                       { f.playing = true }
func (f *fakeVoice) Pause()                      { f.playing = false }
func (f *fakeVoice) Playing() bool               { return f.playing }
func (f *fakeVoice) Position() time.Duration     { return f.pos }
func (f *fakeVoice) Duration() time.Duration     { return f.dur }
func (f *fakeVoice) Seek(at time.Duration) error { f.pos = at; return nil }
func (f *fakeVoice) SetVolume(volume float64)    { f.volume = volume }
func (f *fakeVoice) Close() error                { f.closed = true; return nil }

// fakeResolver answers with one fixed clip or error.
type fakeResolver struct {
	clip Clip
	err  error
}

func (r *fakeResolver) Resolve(context.Context, string, int) (Clip, error) {
	return r.clip, r.err
}

// hangResolver blocks until the resolve context is done.
type hangResolver struct{}

func (hangResolver) Resolve(ctx context.Context, _ string, _ int) (Clip, error) {
	<-ctx.Done()

	return Clip{}, ctx.Err()
}
