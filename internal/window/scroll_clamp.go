package window

// pullScroll clamps the scroll offset to the content the latest relayout
// produced, so a page that got shorter cannot show empty space below it.
func (s *shell) pullScroll() {
	if s.stretched() {
		return
	}

	contentW, contentH := s.contentSize()
	s.scrollX, s.scrollY = clampScroll(s.scrollX, s.scrollY, contentW, contentH, s.screenW, s.screenH)
}
