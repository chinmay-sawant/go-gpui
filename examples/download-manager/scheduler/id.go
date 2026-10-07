package scheduler

import (
	"crypto/rand"
	"encoding/hex"
)

// newID returns a stable, collision-safe job ID.
func newID() string {
	var b [8]byte
	_, _ = rand.Read(b[:])

	return "job-" + hex.EncodeToString(b[:])
}

// signal wakes one waiting worker without blocking.
func (e *Engine) signal() {
	select {
	case e.wake <- struct{}{}:
	default:
	}
}
