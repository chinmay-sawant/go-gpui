package game

import "fmt"

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
func (r Replay) Validate() error {
	if r.Ruleset != "" && r.Ruleset != Ruleset {
		return fmt.Errorf("game: replay ruleset %q, want %q", r.Ruleset, Ruleset)
	}

	if r.FixtureVersion != 0 && r.FixtureVersion != FixtureVersion {
		return fmt.Errorf("game: replay fixture version %d", r.FixtureVersion)
	}

	if r.Fixture != "" && !isFixture(r.Fixture) {
		return fmt.Errorf("game: replay fixture %q", r.Fixture)
	}

	var last uint64
	for i, e := range r.Events {
		if !e.Action.valid() {
			return fmt.Errorf("game: replay event %d has action %d", i, e.Action)
		}

		if i > 0 && e.Step < last {
			return fmt.Errorf("game: replay event %d goes backwards", i)
		}

		last = e.Step
	}

	return nil
}

// isFixture reports whether name is one of FixtureNames.
func isFixture(name string) bool {
	for _, f := range FixtureNames() {
		if f == name {
			return true
		}
	}

	return false
}
