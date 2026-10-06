package scene

import "context"

// pageHistory moves to another live score page.
func (s *Scene) pageHistory(ctx context.Context, delta int) error {
	if !s.hist.open || s.hist.demo {
		return nil
	}

	if delta < 0 && !s.hist.canPrev() {
		return nil
	}

	if delta > 0 && !s.hist.canNext() {
		return nil
	}

	s.hist.page = s.hist.want(s.hist.page + delta)
	s.hist.loading = true
	s.hist.msg = "LOADING"
	s.scoresReq = s.requestScores()

	return s.redraw(ctx)
}

// toggleDemo switches the history screen between live rankings and the
// seeded demo entries.
func (s *Scene) toggleDemo(ctx context.Context) error {
	if !s.hist.open {
		return nil
	}

	s.hist.demo = !s.hist.demo
	s.hist.page = 0
	s.hist.more = false
	s.hist.entries = nil
	s.hist.loading = true
	s.hist.msg = "LOADING"
	s.dirty = true

	if s.hist.demo {
		s.demoReq = s.requestDemo()
	} else {
		s.scoresReq = s.requestScores()
	}

	s.page.SetData(s.view())

	return s.redraw(ctx)
}
