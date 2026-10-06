package host

// ViewportPinner is a screen whose top layers stay at the viewport while
// the page scrolls. The window draws operations whose z-index is at least
// ViewportPinZ() at the offset the display was built with, so a pinned bar
// needs no redraw on scroll.
type ViewportPinner interface {
	ViewportPinZ() int
}
