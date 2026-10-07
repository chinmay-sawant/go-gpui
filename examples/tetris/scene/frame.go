package scene

import "github.com/chinmay-sawant/ownframe/examples/tetris/game"

// Frame is everything the scene draws for one frame. It is a value, so
// the scene compares frames and repaints only when one changed.
type Frame struct {
	Board [game.Rows][game.Cols]game.Cell

	// Piece and Cells are the falling piece; a zero Piece means none.
	Piece game.Piece
	Cells [4]game.Point
	Ghost [4]game.Point
	// Next is the first preview piece.
	Next game.Piece

	Score int
	Lines int
	Level int

	Phase game.Phase
}

// cellKind returns the piece drawn in a board cell and whether the cell
// shows the ghost. Locked cells win, then the falling piece, then the
// ghost on the cells it would land on.
func (f Frame) cellKind(r, c int) (game.Piece, bool) {
	if k := game.Piece(f.Board[r][c]); k != 0 {
		return k, false
	}

	if f.Piece != 0 {
		for _, p := range f.Cells {
			if p.X == c && p.Y == r {
				return f.Piece, false
			}
		}

		for _, p := range f.Ghost {
			if p.X == c && p.Y == r {
				return f.Piece, true
			}
		}
	}

	return 0, false
}

// letter is the lowercase class letter for a piece, or 0.
func letter(k game.Piece) byte {
	if !k.Valid() {
		return 0
	}

	return "iotszjl"[k-1]
}
