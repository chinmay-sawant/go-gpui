package game

// Apply applies one action and reports what it caused. Pause, restart,
// and start work in any phase; the rest need a running game.
func (g *Game) Apply(a Action) []Event {
	switch a {
	case ActionStart:
		return g.Start()
	case ActionPause:
		return g.TogglePause()
	case ActionRestart:
		return g.Restart()
	}

	if g.Phase != PhaseRunning || !g.Piece.Valid() {
		return nil
	}

	switch a {
	case ActionLeft:
		g.move(-1, 0)
	case ActionRight:
		g.move(1, 0)
	case ActionRotateCW:
		g.rotate(true)
	case ActionRotateCCW:
		g.rotate(false)
	case ActionSoftDrop:
		if g.move(0, 1) {
			g.Score = addScore(g.Score, 1)
		}
	case ActionHardDrop:
		return g.hardDrop()
	}

	return nil
}

// String returns the action name.
func (a Action) String() string {
	names := [...]string{
		"none", "left", "right", "rotate-cw", "rotate-ccw", "soft-drop",
		"hard-drop", "pause", "restart", "start",
	}
	if int(a) < len(names) {
		return names[a]
	}

	return "unknown"
}
