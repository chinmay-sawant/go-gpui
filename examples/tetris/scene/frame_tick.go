package scene

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

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

		s.frame = f
		s.dirty = true
	}

	s.watchRun(f)
	s.saveCompleted()
	s.scheduleSave()

	return s.paint(ctx)
}

// watchRun drops the resume slot when a restart begins a new run.
func (s *Scene) watchRun(f Frame) {
	id := s.model.RunID()
	if id == s.runID {
		return
	}

	s.runID = id

	if f.Phase == game.PhaseRunning && f.Score == 0 && f.Lines == 0 {
		s.clearResume()
	}
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
