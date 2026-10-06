package scene

import (
	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
	"github.com/chinmay-sawant/ownframe/examples/tetris/input"
	"github.com/chinmay-sawant/ownframe/examples/tetris/store"
)

// Configure rebuilds the tracker with the saved keymap.
func (c *Core) Configure(s Settings) {
	set, ok := s.Control.(store.Settings)
	if !ok {
		return
	}

	c.keys = input.NewTracker(set.Keymap)
}

// Restore replaces the game with a resumable snapshot.
func (c *Core) Restore(s game.Snapshot) error {
	g, err := game.FromSnapshot(s)
	if err != nil {
		return err
	}

	c.g = g
	c.run = g.ID
	c.rec = nil
	c.done, c.rep, c.taken = nil, nil, false

	return nil
}
