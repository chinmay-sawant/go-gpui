package game

// InputEvent is one action recorded at a simulation step.
type InputEvent struct {
	Step   uint64 `json:"step"`
	Action Action `json:"action"`
}

// Replay is a seed and the input events for a reproducible run.
type Replay struct {
	Seed           uint64       `json:"seed"`
	Ruleset        string       `json:"ruleset"`
	FixtureVersion int          `json:"fixture_version"`
	Fixture        string       `json:"fixture,omitempty"`
	Events         []InputEvent `json:"events"`
}

// Validate rejects a replay this ruleset cannot run.
func (r Replay) Validate() error { return nil }

// Play runs the replay to game over or maxSteps and returns the result.
func (r Replay) Play(maxSteps int) (Result, error) { return Result{}, nil }

// Sequence returns n generated pieces for a seed, for previews and tests.
func Sequence(seed uint64, n int) []Piece { return nil }

// Fixture is a named starting position for tests and demos.
type Fixture struct {
	Name  string
	Board Board
	Piece Piece
	Rot   Rotation
	X, Y  int
	Note  string
}

// Fixtures returns the selectable board fixtures in UI order.
func Fixtures() []Fixture { return nil }

// FixtureNames lists the fixture names.
func FixtureNames() []string { return nil }

// NewFromFixture starts a running game on a named fixture.
func NewFromFixture(name string, seed uint64) (*Game, error) { return nil, nil }
