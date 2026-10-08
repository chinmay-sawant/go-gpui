package host

// ScrollGate is a screen that can refuse page scrolling.
// The window still pans a screen that does not implement it.
type ScrollGate interface {
	AllowScroll() bool
}
