// Package input turns ownframe key events into game actions. ownframe
// forwards one pair per real key event and drops OS auto-repeat pulses, so
// the tracker implements held-key repeat itself: an action fires on press,
// waits DAS, then repeats every ARR. Step emits the actions for one fixed
// simulation step.
package input

import (
	"time"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

// Repeat timing for held keys.
const (
	DAS           = 170 * time.Millisecond
	ARR           = 50 * time.Millisecond
	SoftDropDelay = 40 * time.Millisecond
)

// Keymap maps lowercase key names to game actions; it is stored in
// settings, so the fields carry JSON tags.
type Keymap struct {
	Left      []string `json:"left"`
	Right     []string `json:"right"`
	SoftDrop  []string `json:"soft_drop"`
	HardDrop  []string `json:"hard_drop"`
	RotateCW  []string `json:"rotate_cw"`
	RotateCCW []string `json:"rotate_ccw"`
	Pause     []string `json:"pause"`
	Restart   []string `json:"restart"`
	Start     []string `json:"start"`
}

// DefaultKeymap returns the arrow/WASD layout.
func DefaultKeymap() Keymap { return Keymap{} }

// Tracker owns held-key state and repeat timers.
type Tracker struct{}

// NewTracker returns a tracker using the keymap.
func NewTracker(km Keymap) *Tracker { return &Tracker{} }

// KeyDown records a press and queues its first action.
func (t *Tracker) KeyDown(key string) {}

// KeyUp records a release; the other held direction resumes.
func (t *Tracker) KeyUp(key string) {}

// ReleaseAll clears every held key and timer, for focus loss.
func (t *Tracker) ReleaseAll() {}

// Step advances repeat timers and returns this step's actions.
func (t *Tracker) Step(step time.Duration) []game.Action { return nil }

// Held lists the currently held key names, sorted.
func (t *Tracker) Held() []string { return nil }
