package scene

import (
	"context"
	"time"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

// Tick is the frame callback: it runs the fixed simulation steps and
// repaints the scene. It runs once per window frame before the draw.
func (s *Scene) Tick(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	now := s.now()
	gap := time.Duration(0)

	if !s.last.IsZero() {
		gap = now.Sub(s.last)
		if gap < 0 {
			gap = 0
		}
	}

	s.last = now

	if !s.hasFocus() {
		if s.focused {
			s.focused = false
			s.model.ClearInput()
			s.step.Reset(now)
			s.pauseIfRunning()
		}

		return s.frameTick(ctx)
	}

	if !s.focused {
		s.focused = true
		s.step.Reset(now)

		if s.auto {
			s.auto = false

			if s.frame.Phase == game.PhasePaused {
				s.model.Resume()
			}
		}
	}

	if gap >= game.StallLimit {
		s.step.Reset(now)
		s.pauseIfRunning()
	}

	steps := s.step.Advance(now)
	for i := 0; i < steps; i++ {
		s.model.Step(game.FixedStep)
	}

	return s.frameTick(ctx)
}
