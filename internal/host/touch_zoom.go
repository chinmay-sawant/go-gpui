package host

// TouchZoomPolicy reports whether a screen accepts two-finger touch zoom.
// Hosts that do not implement it keep the default touch zoom behavior.
type TouchZoomPolicy interface {
	TouchZoomAllowed() bool
}
