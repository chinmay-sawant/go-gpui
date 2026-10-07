package game

// FromSnapshot rebuilds a game and returns it paused, so a resumed game
// waits for the player. Call Start to run it.
func FromSnapshot(s Snapshot) (*Game, error) {
	if err := s.Validate(); err != nil {
		return nil, err
	}

	board, err := ParseBoard(s.Board)
	if err != nil {
		return nil, err
	}

	g := &Game{
		ID:      s.ID,
		Seed:    s.Seed,
		Phase:   PhasePaused,
		Board:   board,
		Piece:   s.Piece,
		Rot:     s.Rot,
		X:       s.X,
		Y:       s.Y,
		Next:    append([]Piece(nil), s.Next...),
		bag:     append([]Piece(nil), s.Bag...),
		Score:   s.Score,
		Lines:   s.Lines,
		Level:   s.Level,
		Pieces:  s.Pieces,
		Elapsed: duration(s.ElapsedMS),
		rng:     rng{state: s.RNG},
		fall:    duration(s.FallMS),
		lock:    duration(s.LockMS),
	}

	g.fillPreview()
	g.grounded = !g.canDrop()

	return g, nil
}
