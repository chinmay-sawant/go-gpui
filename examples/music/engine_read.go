package music

import "time"

// Position is the current voice position, or zero.
func (e *Engine) Position() time.Duration {
	if e.voice == nil {
		return 0
	}

	return e.voice.Position()
}

// Duration is the current voice length, or zero.
func (e *Engine) Duration() time.Duration {
	if e.voice == nil {
		return 0
	}

	return e.voice.Duration()
}

// Credit is the free clip behind the current voice.
func (e *Engine) Credit() Clip { return e.clip }

// Err is the last resolve or decode failure.
func (e *Engine) Err() error { return e.err }

// Close releases the current voice.
func (e *Engine) Close() {
	if e.voice != nil {
		_ = e.voice.Close()
		e.voice = nil
	}
}
