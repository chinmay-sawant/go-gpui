package game

// Rune returns the piece letter.
func (p Piece) Rune() byte { return '?' }

// String returns the piece letter.
func (p Piece) String() string { return string(p.Rune()) }

// Valid reports whether p is one of the seven kinds.
func (p Piece) Valid() bool { return p >= PieceI && p <= PieceL }

// CW returns the next rotation clockwise.
func (r Rotation) CW() Rotation { return r }

// CCW returns the next rotation counter-clockwise.
func (r Rotation) CCW() Rotation { return r }

// String returns the rotation name.
func (r Rotation) String() string { return "0" }

// String returns the phase name.
func (p Phase) String() string { return "" }

// String returns the action name.
func (a Action) String() string { return "" }

// String returns the event name.
func (e Event) String() string { return "" }
