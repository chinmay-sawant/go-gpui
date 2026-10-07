package scene

import (
	"context"

	"github.com/chinmay-sawant/ownframe"
	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

// togglePause flips the paused state of a running or paused game.
func (s *Scene) togglePause() {
	switch s.frame.Phase {
	case game.PhaseRunning:
		s.model.Pause()
	case game.PhasePaused:
		s.model.Resume()
	}
}

// toggleTheme switches the stylesheet and saves the choice.
func (s *Scene) toggleTheme(ctx context.Context) error {
	s.dark = !s.dark

	if err := s.page.SetTheme(themeCSS(s.dark)); err != nil {
		s.dark = !s.dark

		return err
	}

	s.themeReq = s.requestTheme()
	s.setStatus("THEME SAVING")

	return s.redraw(ctx)
}

// requestTheme asks the store to persist the theme.
func (s *Scene) requestTheme() uint64 {
	if s.store == nil {
		return 0
	}

	return s.store.SaveTheme(s.dark)
}

// onClick routes a click by the element id. Every click redraws the page
// afterwards, so a handler only changes state.
func (s *Scene) onClick(ctx context.Context, box ownframe.Box) error {
	switch box.ID {
	case "btn-pause":
		s.togglePause()
	case "btn-restart":
		s.model.Restart()
	case "btn-theme", "t-theme":
		return s.toggleTheme(ctx)
	case "btn-scores":
		return s.openHistory(ctx)
	case "btn-prev":
		return s.pageHistory(ctx, -1)
	case "btn-next":
		return s.pageHistory(ctx, +1)
	case "btn-demo":
		return s.toggleDemo(ctx)
	case "btn-close":
		return s.closeHistory(ctx)
	}

	return nil
}
