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

// frameTick diffs the frame, applies worker results, and paints.
func (s *Scene) frameTick(ctx context.Context) error {
	if err := s.drain(ctx); err != nil {
		return err
	}

	f := s.model.Frame()
	if f != s.frame {
		if f.Phase != s.frame.Phase {
			s.overlays = true
		}

		if s.frame.Phase == game.PhaseRunning && f.Phase == game.PhasePaused {
			s.savePoint()
		}

		if f.Phase == game.PhaseRunning && f.Score == 0 && f.Lines == 0 &&
			s.frame.Phase != game.PhaseRunning {
			s.clearResume()
		}

		s.frame = f
		s.dirty = true
	}

	s.saveCompleted()
	s.scheduleSave()

	return s.paint(ctx)
}

// pauseIfRunning pauses a running game and remembers it was automatic,
// so regaining focus can resume it.
func (s *Scene) pauseIfRunning() {
	if s.frame.Phase == game.PhaseRunning {
		s.auto = true
		s.model.Pause()
	}
}

// hasFocus reports the window focus; without a check the scene is focused.
func (s *Scene) hasFocus() bool {
	return s.focus == nil || s.focus()
}
