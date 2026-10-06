package game

// collides reports whether the piece at the box origin overlaps a wall,
// the floor, or a filled cell. Cells above the board are allowed.
func (g *Game) collides(p Piece, r Rotation, x, y int) bool {
	for _, c := range p.Cells(r) {
		bx, by := x+c.X, y+c.Y
		if bx < 0 || bx >= Cols || by >= Rows {
			return true
		}

		if by >= 0 && g.Board[by][bx] != 0 {
			return true
		}
	}

	return false
}

// canDrop reports whether the piece can move down one row.
func (g *Game) canDrop() bool {
	return !g.collides(g.Piece, g.Rot, g.X, g.Y+1)
}

// landY returns the row the piece would rest at.
func (g *Game) landY() int {
	y := g.Y
	for !g.collides(g.Piece, g.Rot, g.X, y+1) {
		y++
	}

	return y
}
