package scene

import "github.com/chinmay-sawant/ownframe/examples/tetris/game"

// watch resets the recording on a new run and captures a finished one.
func (c *Core) watch() {
	if c.g.ID != c.run {
		c.run = c.g.ID
		c.rec = nil
		c.done, c.rep, c.taken = nil, nil, false
	}

	if c.g.Phase != game.PhaseOver || c.taken {
		return
	}

	c.taken = true
	res := c.g.Result()
	c.done = &res
	c.rep = &game.Replay{
		Seed:           c.g.Seed,
		Ruleset:        game.Ruleset,
		FixtureVersion: game.FixtureVersion,
		Events:         append([]game.InputEvent(nil), c.rec...),
	}
}

// TakeCompleted returns a finished run once, with its replay.
func (c *Core) TakeCompleted() (game.Result, *game.Replay, bool) {
	if c.done == nil {
		return game.Result{}, nil, false
	}

	res, rep := *c.done, c.rep
	c.done, c.rep = nil, nil

	return res, rep, true
}

// SavePoint returns a snapshot when the run can resume later.
func (c *Core) SavePoint() (game.Snapshot, bool) {
	switch c.g.Phase {
	case game.PhaseRunning, game.PhasePaused:
		return c.g.Snapshot(), true
	}

	return game.Snapshot{}, false
}
