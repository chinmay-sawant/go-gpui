package window

// devSync follows the page's flag once per frame, so Config.DevTools and a
// SetDevTools call open or close the overlay without a key.
func (s *shell) devSync() error {
	insp := s.devInspector()
	if insp == nil {
		s.dev.on = false

		return nil
	}

	on := insp.DevTools()
	if on == s.dev.on {
		return nil
	}

	s.dev.on = on
	if !on {
		return s.devClose()
	}

	return nil
}
