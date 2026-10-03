package player

import (
	"context"
	"math/rand/v2"
)

// Tick advances audio and the frame animation once per window frame.
func (a *App) Tick(ctx context.Context) error {
	if a.audio == nil {
		return nil
	}

	a.audio.Update()
	a.autoStart(ctx)

	if err := a.advance(ctx); err != nil {
		return err
	}

	a.animate()

	return nil
}

// autoStart requests the current track when playing is wanted but the
// engine has not been asked for this request key yet.
func (a *App) autoStart(ctx context.Context) {
	if !a.view.Playing || a.audio.Loading() || a.audio.Err() != nil {
		return
	}

	if a.audioKey == a.requestKey() {
		return
	}

	a.playNow(ctx)
}

// advance moves on when the voice finishes: repeat rewinds it, shuffle
// picks a random row, and otherwise the queue steps forward.
func (a *App) advance(ctx context.Context) error {
	if !a.view.Playing || a.audio.Loading() || !a.audio.Ended() {
		return nil
	}

	if a.view.Repeat {
		a.audio.SeekFraction(0)
		a.audio.Play()

		return nil
	}

	if a.view.Shuffle && len(a.view.Queue) > 0 {
		a.selectQueue(rand.IntN(len(a.view.Queue)))
	} else {
		a.step(1)
	}

	a.playNow(ctx)
	a.page.SetData(a.view)

	return a.page.Redraw(ctx)
}
