package game

// MaxPlaySteps bounds Replay.Play when the caller passes no limit.
const MaxPlaySteps = 1_000_000

// Play runs the replay to game over or maxSteps and returns the final
// result. The score is reproducible for the same seed, fixture, ruleset,
// and events.
func (r Replay) Play(maxSteps int) (Result, error) {
	if err := r.Validate(); err != nil {
		return Result{}, err
	}

	if maxSteps <= 0 {
		maxSteps = MaxPlaySteps
	}

	var g *Game
	if r.Fixture != "" {
		var err error

		g, err = NewFromFixture(r.Fixture, r.Seed)
		if err != nil {
			return Result{}, err
		}
	} else {
		g = New(r.Seed)
		g.Start()
	}

	i := 0
	for step := 0; step < maxSteps && g.Phase != PhaseOver; step++ {
		for i < len(r.Events) && r.Events[i].Step == uint64(step) {
			g.Apply(r.Events[i].Action)
			i++
		}

		g.Step(FixedStep)
	}

	return g.Result(), nil
}
