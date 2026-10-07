package window

// frameMarker is a ticking page that can opt into the partial buffer.
type frameMarker interface {
	FrameDirty() bool
}

// fullReplay reports a ticking page that did not opt into a dirty rect.
// UseFrameDirty keeps that page on the buffer path.
func (s *shell) fullReplay() bool {
	if !s.isTicking() {
		return false
	}

	marker, ok := s.app.(frameMarker)

	return !ok || !marker.FrameDirty()
}
