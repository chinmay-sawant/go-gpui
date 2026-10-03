package music

import "time"

// replace closes the old voice and installs a new one.
func (e *Engine) replace(voice Voice, clip Clip) {
	if e.voice != nil {
		_ = e.voice.Close()
	}

	voice.SetVolume(e.volume)
	e.voice = voice
	e.clip = clip

	if e.wantPlay {
		voice.Play()
	}
}

// Play resumes the current voice, or marks the pending one to start.
func (e *Engine) Play() {
	e.wantPlay = true
	if e.voice != nil {
		e.voice.Play()
	}
}

// Pause stops the current voice and holds a pending one.
func (e *Engine) Pause() {
	e.wantPlay = false
	if e.voice != nil {
		e.voice.Pause()
	}
}

// Toggle flips between playing and paused.
func (e *Engine) Toggle() {
	if e.Playing() {
		e.Pause()

		return
	}

	e.Play()
}

// Playing reports audio that is actually advancing.
func (e *Engine) Playing() bool {
	return e.voice != nil && e.wantPlay && e.voice.Playing()
}

// Ended reports a finished voice, so the app can advance to the next track.
func (e *Engine) Ended() bool {
	return e.voice != nil && !e.Playing() && e.Duration() > 0 && e.Position() >= e.Duration()
}

// Loading reports a resolve still in flight.
func (e *Engine) Loading() bool { return e.loading }

// SeekFraction moves inside the current voice by a 0..1 fraction.
func (e *Engine) SeekFraction(fraction float64) {
	if e.voice == nil {
		return
	}

	if fraction < 0 {
		fraction = 0
	}

	if fraction > 1 {
		fraction = 1
	}

	_ = e.voice.Seek(time.Duration(fraction * float64(e.Duration())))
}

// SetVolume sets the engine volume from 0 to 1.
func (e *Engine) SetVolume(volume float64) {
	e.volume = clampVolume(volume)
	if e.voice != nil {
		e.voice.SetVolume(e.volume)
	}
}
