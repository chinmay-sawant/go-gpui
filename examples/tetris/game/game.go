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

	rng rng

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

// Start spawns the first piece and begins play.
func (g *Game) Start() {}

// Restart resets every transient value and begins a new run.
func (g *Game) Restart() {}

// TogglePause pauses a running game and resumes a paused one.
func (g *Game) TogglePause() {}

// SetPaused sets the paused state.
func (g *Game) SetPaused(paused bool) {}

// Apply applies one action and reports what it caused.
func (g *Game) Apply(a Action) []Event { return nil }

// Step advances the simulation by one fixed step and reports events.
func (g *Game) Step(step time.Duration) []Event { return nil }

// ActiveCells returns the current piece's absolute board cells.
func (g *Game) ActiveCells() []Point { return nil }

// GhostCells returns where the current piece would land.
func (g *Game) GhostCells() []Point { return nil }

// NextPiece returns the first preview piece.
func (g *Game) NextPiece() Piece { return PieceI }

// Result summarizes the run for storage.
func (g *Game) Result() Result { return Result{} }

// Snapshot captures the resumable state.
func (g *Game) Snapshot() Snapshot { return Snapshot{} }
