package host

import "time"

// Stats is one snapshot of the page counters an inspector shows. The zero
// value is a page that has drawn nothing.
type Stats struct {
	Redraws    uint64
	Parses     uint64
	Cascades   uint64
	Layouts    uint64
	Repaints   uint64
	Boxes      int
	Ops        int
	LastRedraw time.Duration
	LastDraw   time.Duration
	Reloads    uint64
	// LastReloadError is the most recent failed hot reload, empty when none.
	LastReloadError string
}

// Inspector is a screen with a devtools overlay. The window draws the
// overlay; the page reports the stats because only the page can measure a
// parse.
type Inspector interface {
	DevTools() bool
	SetDevTools(on bool)
	Stats() Stats
}
