package game

import (
	"errors"
	"time"
)

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
	Bag            []Piece  `json:"bag"`
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

// Snapshot captures the resumable state. Call it between steps, never
// inside one.
func (g *Game) Snapshot() Snapshot {
	return Snapshot{
		ID:             g.ID,
		Board:          g.Board.rows(),
		Piece:          g.Piece,
		Rot:            g.Rot,
		X:              g.X,
		Y:              g.Y,
		Next:           append([]Piece(nil), g.Next...),
		Bag:            append([]Piece(nil), g.bag...),
		Score:          g.Score,
		Lines:          g.Lines,
		Level:          g.Level,
		Pieces:         g.Pieces,
		ElapsedMS:      g.Elapsed.Milliseconds(),
		Seed:           g.Seed,
		RNG:            g.rng.state,
		FallMS:         g.fall.Milliseconds(),
		LockMS:         g.lock.Milliseconds(),
		Ruleset:        Ruleset,
		FixtureVersion: FixtureVersion,
	}
}

// duration converts stored milliseconds to a Duration.
func duration(ms int64) time.Duration {
	return time.Duration(ms) * time.Millisecond
}
