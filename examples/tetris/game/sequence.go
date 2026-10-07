package game

// Sequence returns n generated pieces for a seed. Two calls with the same
// seed return the same pieces, so tests and demos can pin a sequence.
func Sequence(seed uint64, n int) []Piece {
	if n < 0 {
		n = 0
	}

	g := New(seed)
	out := make([]Piece, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, g.nextPiece())
	}

	return out
}
