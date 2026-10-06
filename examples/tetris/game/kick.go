package game

// kick is one wall-kick offset in board cells.
type kick struct{ dx, dy int }

// kicks returns the SRS offsets to try for a rotation, in order. The
// tables are the standard Super Rotation System data with y flipped,
// because board rows grow downward.
func kicks(p Piece, from Rotation, cw bool) [5]kick {
	if p == PieceO {
		return [5]kick{}
	}

	dir := 0
	if !cw {
		dir = 1
	}

	if p == PieceI {
		return iKicks[from%4][dir]
	}

	return jlstzKicks[from%4][dir]
}

// jlstzKicks holds the T, S, Z, J, and L kick offsets.
var jlstzKicks = [4][2][5]kick{
	{
		{{0, 0}, {-1, 0}, {-1, -1}, {0, 2}, {-1, 2}},
		{{0, 0}, {1, 0}, {1, -1}, {0, 2}, {1, 2}},
	},
	{
		{{0, 0}, {1, 0}, {1, 1}, {0, -2}, {1, -2}},
		{{0, 0}, {1, 0}, {1, 1}, {0, -2}, {1, -2}},
	},
	{
		{{0, 0}, {1, 0}, {1, -1}, {0, 2}, {1, 2}},
		{{0, 0}, {-1, 0}, {-1, -1}, {0, 2}, {-1, 2}},
	},
	{
		{{0, 0}, {-1, 0}, {-1, 1}, {0, -2}, {-1, -2}},
		{{0, 0}, {-1, 0}, {-1, 1}, {0, -2}, {-1, -2}},
	},
}

// iKicks holds the I piece kick offsets.
var iKicks = [4][2][5]kick{
	{
		{{0, 0}, {-2, 0}, {1, 0}, {-2, 1}, {1, -2}},
		{{0, 0}, {-1, 0}, {2, 0}, {-1, -2}, {2, 1}},
	},
	{
		{{0, 0}, {-1, 0}, {2, 0}, {-1, -2}, {2, 1}},
		{{0, 0}, {2, 0}, {-1, 0}, {2, -1}, {-1, 2}},
	},
	{
		{{0, 0}, {2, 0}, {-1, 0}, {2, -1}, {-1, 2}},
		{{0, 0}, {1, 0}, {-2, 0}, {1, 2}, {-2, -1}},
	},
	{
		{{0, 0}, {1, 0}, {-2, 0}, {1, 2}, {-2, -1}},
		{{0, 0}, {-2, 0}, {1, 0}, {-2, 1}, {1, -2}},
	},
}
