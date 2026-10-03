package music

import (
	"sync"
	"time"
)

// Result is one resolved request waiting for the frame thread.
type Result struct {
	Gen  uint64
	Clip Clip
	Err  error
}

// Engine resolves one track at a time off the frame thread and drives one
// decoded Voice on it. An example calls Want from a click and Update from
// the page tick, then reads Position and Credit for the moving UI.
type Engine struct {
	// Open turns encoded bytes into a Voice. Tests replace it with a fake.
	Open func(data []byte) (Voice, error)

	// Timeout caps one background resolve, so a hung request cannot leave
	// the UI loading forever. Zero means defaultTimeout.
	Timeout time.Duration

	res      Resolver
	latestMu sync.Mutex
	latest   Result

	gen      uint64
	applied  uint64
	loading  bool
	clip     Clip
	voice    Voice
	wantPlay bool
	volume   float64
	err      error
}

// NewEngine returns an engine over res, playing through NewSound.
func NewEngine(res Resolver) *Engine {
	return &Engine{
		Open:    func(data []byte) (Voice, error) { return NewSound(data) },
		Timeout: defaultTimeout,
		res:     res,
		volume:  1,
	}
}

// defaultTimeout is the resolve cap when Timeout is zero.
const defaultTimeout = 20 * time.Second

// Update applies a resolved track and starts it when wanted. Call it from the
// page tick; it never blocks.
func (e *Engine) Update() {
	e.latestMu.Lock()
	res := e.latest
	e.latestMu.Unlock()

	if !e.loading || res.Gen <= e.applied {
		return
	}

	e.applied = res.Gen
	e.loading = false

	if res.Err != nil {
		e.err = res.Err

		return
	}

	voice, err := e.Open(res.Clip.Data)
	if err != nil {
		e.err = err

		return
	}

	e.replace(voice, res.Clip)
}
