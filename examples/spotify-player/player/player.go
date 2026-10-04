// Package player is the Spotify-like player example.
// The screen is one HTML template assembled from per-component fragments in
// components/. The sidebar, greeting, shelf, tracklist, and now bar read
// their strings from View. A click selects a track, a card, or a playlist.
// Load replaces the sample data with live iTunes search results. gpui opens
// the window. This package does not.
package player

const (
	// DefaultWidth and DefaultHeight are the size of a newly opened window.
	DefaultWidth  = 1440
	DefaultHeight = 960

	// MinWidth and MinHeight are the smallest frame the screen will draw.
	MinWidth  = 1120
	MinHeight = 720

	// MaxWidth and MaxHeight cap a resized frame.
	MaxWidth  = 2560
	MaxHeight = 2560

	// defaultVolume is the volume a fresh view starts at, in percent.
	defaultVolume = 10
)
