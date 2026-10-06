package game

// Cells returns the piece's cells in its own box for the rotation.
func (p Piece) Cells(r Rotation) []Point { return nil }

// Filled reports whether the board cell is occupied.
func (b Board) Filled(y, x int) bool { return false }

// ParseBoard reads rows of text: '.' empty, piece letters filled.
func ParseBoard(rows []string) (Board, error) { return Board{}, nil }
