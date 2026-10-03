package player

import (
	"context"
	"strconv"
	"strings"
)

// query is the search text for the current track.
func (a *App) query() string {
	if q := strings.TrimSpace(a.view.Now.Genre); q != "" {
		return q
	}

	return "instrumental"
}

// requestKey names the audio request behind the current track.
func (a *App) requestKey() string {
	return a.query() + "#" + strconv.Itoa(a.view.Now.Index)
}

// playNow asks the engine for the current track and shows playing intent.
func (a *App) playNow(ctx context.Context) {
	if a.audio == nil {
		return
	}

	a.audioKey = a.requestKey()
	a.audio.Want(ctx, a.query(), a.view.Now.Index)
	a.view.Playing = true
}

// togglePlay flips the play intent and starts, resumes, or pauses the voice.
func (a *App) togglePlay(ctx context.Context) {
	a.view.Playing = !a.view.Playing

	if a.audio == nil {
		return
	}

	switch {
	case !a.view.Playing:
		a.audio.Pause()
	case a.audio.Duration() > 0:
		a.audio.Play()
	default:
		a.playNow(ctx)
	}
}

// seekAudio moves the voice to the view's progress.
func (a *App) seekAudio() {
	if a.audio != nil {
		a.audio.SeekFraction(float64(a.view.Progress) / 100)
	}
}

// volumeAudio copies the view's volume to the engine.
func (a *App) volumeAudio() {
	if a.audio != nil {
		a.audio.SetVolume(float64(a.view.Volume) / 100)
	}
}
