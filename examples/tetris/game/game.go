package game

import "time"

// Game is one game of Tetris. The exported fields are the state a scene
// draws; the unexported fields are timers and generators.
type Game struct {
	ID      string
	Seed    uint64
	Phase   Phase
	Board   Board
	Piece   Piece
	Rot     Rotation
	X, Y    int
	Next    []Piece
	Score   int
	Lines   int
	Level   int
	Pieces  int
	Elapsed time.Duration
	Steps   uint64

	rng      rng
	bag      []Piece
	fall     time.Duration
	lock     time.Duration
	resets   int
	grounded bool
}

// New returns a ready game with an empty board and a seeded piece bag.
func New(seed uint64) *Game {
	return &Game{
		ID:    newID(),
		Seed:  seed,
		Phase: PhaseReady,
		Level: 1,
		rng:   rng{state: seed},
	}
}

// Result summarizes the run for storage.
func (g *Game) Result() Result {
	return Result{
		ID:             g.ID,
		Score:          g.Score,
		Lines:          g.Lines,
		Level:          g.Level,
		Pieces:         g.Pieces,
		DurationMS:     g.Elapsed.Milliseconds(),
		Seed:           g.Seed,
		Ruleset:        Ruleset,
		FixtureVersion: FixtureVersion,
	}
}
