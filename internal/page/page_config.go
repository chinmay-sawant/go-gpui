package page

// Config is the template, the optional theme, and the frame size for a new
// page. Theme is an extra stylesheet applied after the template's own styles;
// empty means no theme.
// Width and Height are the first frame, in CSS pixels.
// MinWidth and MinHeight are the smallest frame. Zero means 1.
// MaxWidth and MaxHeight ask the host to cap the window. Zero means no cap.
// They do not cap the picture, so the layout follows the window.
type Config struct {
	Title string
	HTML  string
	// File reads the page source from disk at New and watches it. Setting
	// File and HTML together is an error.
	File  string
	Theme string
	// ThemeFile reads the theme stylesheet from disk and watches it.
	ThemeFile string
	// DisableHotReload turns the file watch off. It is ignored without File
	// or ThemeFile.
	DisableHotReload bool
	// DevTools starts the window overlay on.
	DevTools bool
	// Perf records pipeline timing, dirty counts, and allocs in Stats. It
	// is off by default, so end users pay nothing; SetPerf opts in.
	Perf bool
	// LockView scales the page to the window and turns off scroll and zoom.
	LockView  bool
	Width     int
	Height    int
	MinWidth  int
	MinHeight int
	MaxWidth  int
	MaxHeight int
}
