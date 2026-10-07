package scene

import "github.com/chinmay-sawant/ownframe/examples/tetris/game"

// Frame maps the game state onto the scene frame.
func (c *Core) Frame() Frame {
	f := Frame{
		Board: c.g.Board,
		Piece: c.g.Piece,
		Score: c.g.Score,
		Lines: c.g.Lines,
		Level: c.g.Level,
		Phase: c.g.Phase,
	}

	if len(c.g.Next) > 0 {
		f.Next = c.g.Next[0]
	}

	if c.g.Piece.Valid() {
		f.Cells = pointsToCells(c.g.ActiveCells())
		f.Ghost = pointsToCells(c.g.GhostCells())
	}

	return f
}

// pointsToCells copies up to four board cells into a fixed array.
func pointsToCells(pts []game.Point) [4]game.Point {
	var out [4]game.Point

	for i := 0; i < len(pts) && i < 4; i++ {
		out[i] = pts[i]
	}

	return out
}
