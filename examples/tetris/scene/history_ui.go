package scene

import (
	"context"

	"github.com/chinmay-sawant/ownframe/examples/tetris/game"
)

// openHistory pauses a running game and loads the score screen. The board
// stays in the model, so closing history restores the same game.
func (s *Scene) openHistory(ctx context.Context) error {
	if s.hist.open {
		return nil
	}

	s.hist.open = true
	s.hist.page = 0
	s.hist.demo = false
	s.hist.more = false
	s.hist.entries = nil
	s.hist.loading = true
	s.hist.msg = "LOADING"

	if s.frame.Phase == game.PhaseRunning {
		s.hist.wasRun = true
		s.model.ClearInput()
		s.model.Pause()
	}

	s.dirty = true
	s.scoresReq = s.requestScores()
	s.page.SetData(s.view())

	return s.page.Load(ctx, pageHTML(historyHTML))
}

// closeHistory returns to the game screen and resumes a game that the
// history screen paused.
func (s *Scene) closeHistory(ctx context.Context) error {
	if !s.hist.open {
		return nil
	}

	s.hist.open = false
	s.dirty = true

	if s.hist.wasRun {
		s.model.Resume()
	}

	s.hist.wasRun = false

	return s.page.Back(ctx)
}

// requestScores asks for the live page; an absent store answers with a
// message instead.
func (s *Scene) requestScores() uint64 {
	if s.store == nil {
		s.hist.loading = false
		s.hist.msg = "SCORES UNAVAILABLE"

		return 0
	}

	return s.store.Scores(s.hist.page)
}

// requestDemo asks for the seeded demo entries.
func (s *Scene) requestDemo() uint64 {
	if s.store == nil {
		s.hist.loading = false
		s.hist.msg = "SCORES UNAVAILABLE"

		return 0
	}

	return s.store.DemoScores()
}
