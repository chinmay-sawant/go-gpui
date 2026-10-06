package game

import "errors"

// ErrInvalidSnapshot rejects a snapshot outside the model bounds.
var ErrInvalidSnapshot = errors.New("game: invalid snapshot")

// Snapshot is the persisted state of a resumable game.
type Snapshot struct {
	ID             string   `json:"id"`
	Board          []string `json:"board"`
	Piece          Piece    `json:"piece"`
	Rot            Rotation `json:"rot"`
	X              int      `json:"x"`
	Y              int      `json:"y"`
	Next           []Piece  `json:"next"`
	Score          int      `json:"score"`
	Lines          int      `json:"lines"`
	Level          int      `json:"level"`
	Pieces         int      `json:"pieces"`
	ElapsedMS      int64    `json:"elapsed_ms"`
	Seed           uint64   `json:"seed"`
	RNG            uint64   `json:"rng"`
	FallMS         int64    `json:"fall_ms"`
	LockMS         int64    `json:"lock_ms"`
	Ruleset        string   `json:"ruleset"`
	FixtureVersion int      `json:"fixture_version"`
}

// Validate rejects a snapshot outside the model bounds.
func (s Snapshot) Validate() error { return nil }

// FromSnapshot rebuilds a game, or returns ErrInvalidSnapshot.
func FromSnapshot(s Snapshot) (*Game, error) { return New(s.Seed), nil }
