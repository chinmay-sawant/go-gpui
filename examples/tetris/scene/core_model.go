package scene

import (
	"time"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
	"github.com/chinmay-sawant/ownframe/examples/tetris/input"
)

// Core is the Model backed by the core game and input tracker. It owns
// replay recording, completion capture, and resume points.
type Core struct {
	g     *game.Game
	keys  *input.Tracker
	run   string
	rec   []game.InputEvent
	done  *game.Result
	rep   *game.Replay
	taken bool
}

// NewCore returns a model on a fresh game.
func NewCore(seed uint64) *Core {
	g := game.New(seed)

	return &Core{g: g, keys: input.NewTracker(input.DefaultKeymap()), run: g.ID}
}

// RunID names the current run.
func (c *Core) RunID() string { return c.g.ID }

// Down forwards a key press to the tracker.
func (c *Core) Down(key string) { c.keys.KeyDown(key) }

// Up forwards a key release to the tracker.
func (c *Core) Up(key string) { c.keys.KeyUp(key) }

// ClearInput drops every held key.
func (c *Core) ClearInput() { c.keys.ReleaseAll() }

// Step runs one fixed step: queued input, the game, then the run watch.
func (c *Core) Step(step time.Duration) {
	for _, a := range c.keys.Step(step) {
		if a >= game.ActionLeft && a <= game.ActionHardDrop {
			c.rec = append(c.rec, game.InputEvent{Step: c.g.Steps, Action: a})
		}

		c.g.Apply(a)
	}

	c.g.Step(step)
	c.watch()
}

// Pause pauses a running game.
func (c *Core) Pause() { c.g.SetPaused(true) }

// Resume resumes a paused game.
func (c *Core) Resume() { c.g.SetPaused(false) }

// Restart begins a fresh run and drops the old recording.
func (c *Core) Restart() {
	c.g.Restart()
	c.watch()
}
