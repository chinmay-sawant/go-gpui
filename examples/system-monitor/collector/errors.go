package collector

import (
	"errors"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/system-monitor/domain"
)

// ErrStarted reports a second Start on one manager.
var ErrStarted = errors.New("system-monitor: collector already started")

// ErrClosed reports work on a closed manager.
var ErrClosed = errors.New("system-monitor: collector closed")

// ErrPoolFull reports a saturated worker queue.
var ErrPoolFull = errors.New("system-monitor: collector queue full")

// Error is the latest failure of one collector step. The UI shows it next to
// the values it could not refresh.
type Error struct {
	Step   string
	Source string
	At     time.Time
	Err    string
}

// Sink receives every published sample while recording is on. RecordSample
// must not block the collector: storage's Recorder copies into a bounded
// queue and counts drops. The return value reports whether the sample was
// queued; the manager ignores it, because a drop is already counted.
type Sink interface {
	RecordSample(sample domain.Sample) bool
}
