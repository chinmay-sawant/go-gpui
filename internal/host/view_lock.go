package host

// ViewLock is a screen that fills the window.
// The window scales it to the screen. It does not scroll or pinch-zoom.
type ViewLock interface {
	ViewLocked() bool
}
