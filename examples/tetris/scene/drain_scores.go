package scene

import "context"

// applyScores installs a score page, unless a newer request replaced it
// or the screen switched between live and demo mode.
func (s *Scene) applyScores(ctx context.Context, res Result) error {
	if res.Kind == ResultScores {
		if res.ID != s.scoresReq || s.hist.demo {
			return nil
		}
	}

	if res.Kind == ResultDummy {
		if res.ID != s.demoReq || !s.hist.demo {
			return nil
		}
	}

	if res.Err != nil {
		s.hist.loading = false
		s.hist.msg = "SCORES UNAVAILABLE"

		if s.hist.open {
			return s.redraw(ctx)
		}

		return nil
	}

	s.hist.apply(res.Scores)

	if s.hist.open {
		return s.redraw(ctx)
	}

	return nil
}
