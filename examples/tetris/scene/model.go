package scene

import (
	"time"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

// Model is the game the scene draws and steps. The core game and input
// tracker sit behind it; tests use a fake.
type Model interface {
	// Frame returns the state to draw.
	Frame() Frame
	// RunID names the current run; a change means a restart.
	RunID() string
	// Down and Up forward one key event to the input adapter.
	Down(key string)
	Up(key string)
	// ClearInput drops every held key, for focus loss.
	ClearInput()
	// Step runs one fixed simulation step: input repeat, actions, game.
	Step(step time.Duration)
	// Pause and Resume change the paused state of a running game.
	Pause()
	Resume()
	// Restart resets every transient value and begins a new run.
	Restart()
	// Configure applies saved control settings such as the keymap.
	Configure(Settings)
	// Restore replaces the game with a resumable snapshot.
	Restore(game.Snapshot) error
	// SavePoint returns a snapshot when the current run can resume.
	SavePoint() (game.Snapshot, bool)
	// TakeCompleted returns a finished run once, for storage.
	TakeCompleted() (game.Result, *game.Replay, bool)
}

// Stepper converts wall time into fixed steps. game.Clock is the core one.
type Stepper interface {
	Advance(now time.Time) int
	Reset(now time.Time)
}

// Settings are the saved preferences the scene understands.
type Settings struct {
	Dark bool
	// Control is the core control configuration (keymap, ghost), opaque
	// to the scene and applied by the Model adapter.
	Control any
}
