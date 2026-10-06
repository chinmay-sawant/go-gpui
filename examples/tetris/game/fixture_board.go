package game

// nearTopOutBoard leaves two clear rows above a tall stack with a well
// in the middle four columns.
func nearTopOutBoard() Board {
	var b Board
	for y := 2; y < Rows; y++ {
		for x := 0; x < Cols; x++ {
			if x >= 3 && x <= 6 {
				continue
			}

			b[y][x] = Cell(PieceJ)
		}
	}

	return b
}

// wellBoard fills the n bottom rows except column 0, where the fixture I
// piece drops in.
func wellBoard(n int) Board {
	var b Board
	for y := Rows - n; y < Rows; y++ {
		for x := 1; x < Cols; x++ {
			b[y][x] = Cell(PieceL)
		}
	}

	return b
}
