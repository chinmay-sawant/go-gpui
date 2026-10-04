// Package insights is the Wispr Flow insights dashboard example.
// The screen is one HTML template assembled from per-card fragments in
// components/. Every card reads its numbers from View. A click on a tab
// switches the active tab. gpui opens the window. This package does not.
package insights

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 1638
	DefaultHeight = 950

	// MinWidth and MinHeight are the smallest frame the screen will draw.
	MinWidth  = 480
	MinHeight = 560
)
