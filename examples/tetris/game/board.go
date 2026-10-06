package game

import "fmt"

// Filled reports whether the board cell is occupied. Out-of-range cells
// are not filled.
func (b Board) Filled(y, x int) bool {
	return y >= 0 && y < Rows && x >= 0 && x < Cols && b[y][x] != 0
}

// ParseBoard reads rows: '.' empty, piece letters filled.
func ParseBoard(rows []string) (Board, error) {
	var b Board
	if len(rows) != Rows {
		return b, fmt.Errorf("game: board has %d rows, want %d", len(rows), Rows)
	}

	for y, row := range rows {
		if len(row) != Cols {
			return b, fmt.Errorf("game: row %d has %d cells, want %d", y, len(row), Cols)
		}

		for x := 0; x < Cols; x++ {
			if row[x] == '.' {
				continue
			}

			p := pieceFor(row[x])
			if !p.Valid() {
				return b, fmt.Errorf("game: row %d cell %d has %q", y, x, row[x])
			}

			b[y][x] = Cell(p)
		}
	}

	return b, nil
}

// pieceFor maps a letter to a piece, upper or lower case.
func pieceFor(c byte) Piece {
	for p := PieceI; p <= PieceL; p++ {
		if p.Rune() == c || p.Rune()+32 == c {
			return p
		}
	}

	return 0
}

// rows renders the board as text rows that ParseBoard can read back.
func (b Board) rows() []string {
	out := make([]string, Rows)
	for y := range b {
		row := make([]byte, Cols)
		for x := range b[y] {
			row[x] = '.'
			if b[y][x] != 0 {
				row[x] = Piece(b[y][x]).Rune()
			}
		}

		out[y] = string(row)
	}

	return out
}

// fullRows reports which rows are complete.
func (b Board) fullRows() [Rows]bool {
	var full [Rows]bool
	for y := range b {
		full[y] = true

		for x := range b[y] {
			if b[y][x] == 0 {
				full[y] = false

				break
			}
		}
	}

	return full
}

// clearRows removes every full row at once and shifts the rows above
// down. It returns how many rows went away.
func (b *Board) clearRows(full [Rows]bool) int {
	n := 0
	write := Rows - 1

	for y := Rows - 1; y >= 0; y-- {
		if full[y] {
			n++

			continue
		}

		if write != y {
			b[write] = b[y]
		}

		write--
	}

	for ; write >= 0; write-- {
		b[write] = [Cols]Cell{}
	}

	return n
}
