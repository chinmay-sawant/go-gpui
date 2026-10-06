package host

// ScrollObserver is a screen that follows the window scroll offset. The
// window owns the offset; it reports every change here so a virtualized
// page can slice its rows before the next draw. Windowing pages opt into
// a redraw on each scroll change; other pages scroll by blitting.
type ScrollObserver interface {
	SetScrollOffset(x, y int)
	ScrollOffset() (int, int)
	Windowing() bool
}

// ScrollWindowStepper lets a screen reuse an unchanged overscan window.
type ScrollWindowStepper interface {
	StepScrollWindow() bool
}
