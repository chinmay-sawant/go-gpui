package game

import "time"

// LockDelay is how long a grounded piece waits before it locks.
const LockDelay = 500 * time.Millisecond

// maxLockResets caps move and rotate lock-delay resets.
const maxLockResets = 15

// resetLock restarts the lock delay when the piece is grounded, up to the
// reset cap.
func (g *Game) resetLock() {
	if g.canDrop() {
		return
	}

	g.grounded = true
	if g.resets < maxLockResets {
		g.lock = 0
		g.resets++
	}
}

// lockPiece stores the piece, clears full rows at once, and spawns the
// next piece. It tops out when the piece locked above the board.
func (g *Game) lockPiece() []Event {
	ev := []Event{{Kind: EventLock, Score: g.Score}}
	onBoard := false

	for _, c := range g.Piece.Cells(g.Rot) {
		x, y := g.X+c.X, g.Y+c.Y
		if y < 0 {
			continue
		}

		g.Board[y][x] = Cell(g.Piece)
		onBoard = true
	}

	g.Pieces++

	if !onBoard {
		g.topOut()

		return append(ev, Event{Kind: EventTopOut, Score: g.Score})
	}

	full := g.Board.fullRows()
	if n := g.Board.clearRows(full); n > 0 {
		g.Lines = min(g.Lines+n, MaxLines)
		g.Score = addScore(g.Score, lineScore(n)*g.Level)
		ev = append(ev, Event{Kind: EventClear, Lines: n, Score: g.Score, Level: g.Level})

		if lv := levelFor(g.Lines); lv > g.Level {
			g.Level = lv
			ev = append(ev, Event{Kind: EventLevelUp, Level: lv})
		}
	}

	return append(ev, g.spawn()...)
}

// topOut ends the game.
func (g *Game) topOut() {
	g.Phase = PhaseOver
}
