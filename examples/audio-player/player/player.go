// Package player is the Aurora audio player example.
// The screen is one HTML template assembled from per-card fragments in
// components/. The sidebar, now-playing card, recent tiles, and queue read
// their data from View. A click runs one control. Load fills the queue from
// the iTunes search API and falls back to sample data. gpui opens the
// window. This package does not.
package player

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 1280
	DefaultHeight = 860

	// MinWidth and MinHeight are the smallest frame the screen will draw.
	MinWidth  = 1040
	MinHeight = 720

	// MaxWidth and MaxHeight cap a resized frame.
	MaxWidth  = 2560
	MaxHeight = 2560
)
