package scene

import "context"

// applySettings installs the saved theme and control configuration and
// restores a resumable snapshot.
func (s *Scene) applySettings(ctx context.Context, res Result) {
	if res.Err != nil {
		s.setStatus("SETTINGS UNAVAILABLE")

		return
	}

	s.model.Configure(res.Set.Settings)

	if res.Set.Has {
		if err := s.model.Restore(res.Set.Snapshot); err != nil {
			s.setStatus("RESUME REJECTED")
		} else {
			s.setStatus("GAME RESUMED")
		}
	}

	if res.Set.Settings.Dark == s.dark {
		return
	}

	s.dark = res.Set.Settings.Dark

	if err := s.page.SetTheme(themeCSS(s.dark)); err != nil {
		s.setStatus("THEME FAILED")

		return
	}

	if err := s.redraw(ctx); err != nil {
		s.setStatus("THEME FAILED")
	}
}
