// Package input turns ownframe key events into game actions. ownframe
// forwards one pair per real key event and drops OS auto-repeat pulses, so
// the tracker repeats held movement, rotation, and soft drop in Go: the
// action fires on press, waits DAS, then repeats every ARR. Soft drop
// waits SoftDropDelay instead of DAS. Hard drop, pause, restart, and start
// fire once per press. Step emits the actions for one fixed simulation
// step.
package input

import "time"

// Repeat timing for held keys.
const (
	DAS           = 170 * time.Millisecond
	ARR           = 50 * time.Millisecond
	SoftDropDelay = 40 * time.Millisecond
)

// Keymap maps lowercase key names to game actions; it is stored in
// settings, so the fields carry JSON tags. Several keys may map to one
// action.
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
func DefaultKeymap() Keymap {
	return Keymap{
		Left:      []string{"arrowleft", "a"},
		Right:     []string{"arrowright", "d"},
		SoftDrop:  []string{"arrowdown", "s"},
		HardDrop:  []string{"space"},
		RotateCW:  []string{"arrowup", "w", "x"},
		RotateCCW: []string{"z"},
		Pause:     []string{"p", "escape"},
		Restart:   []string{"r"},
		Start:     []string{"enter"},
	}
}
