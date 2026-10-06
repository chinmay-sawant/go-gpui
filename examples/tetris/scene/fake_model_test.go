package scene

import (
	"time"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

// fakeModel records what the scene asked of the game.
type fakeModel struct {
	f Frame

	steps    int
	pauses   int
	resumes  int
	restarts int
	clears   int
	keys     []string
	ups      []string

	configs    []Settings
	restores   []game.Snapshot
	restoreErr error

	snap   game.Snapshot
	snapOK bool
	done   *game.Result
	rep    *game.Replay
}

func (m *fakeModel) Frame() Frame { return m.f }

func (m *fakeModel) Down(key string) { m.keys = append(m.keys, key) }

func (m *fakeModel) Up(key string) { m.ups = append(m.ups, key) }

func (m *fakeModel) ClearInput() { m.clears++ }

func (m *fakeModel) Step(time.Duration) { m.steps++ }

func (m *fakeModel) Pause() {
	m.pauses++

	if m.f.Phase == game.PhaseRunning {
		m.f.Phase = game.PhasePaused
	}
}

func (m *fakeModel) Resume() {
	m.resumes++

	if m.f.Phase == game.PhasePaused {
		m.f.Phase = game.PhaseRunning
	}
}

func (m *fakeModel) Restart() { m.restarts++ }

func (m *fakeModel) Configure(s Settings) { m.configs = append(m.configs, s) }

func (m *fakeModel) Restore(s game.Snapshot) error {
	m.restores = append(m.restores, s)

	return m.restoreErr
}

func (m *fakeModel) SavePoint() (game.Snapshot, bool) { return m.snap, m.snapOK }

func (m *fakeModel) TakeCompleted() (game.Result, *game.Replay, bool) {
	if m.done == nil {
		return game.Result{}, nil, false
	}

	r := *m.done
	m.done = nil

	return r, m.rep, true
}
