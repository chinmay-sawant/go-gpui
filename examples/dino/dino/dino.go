// Package dino shows the Chrome-style dinosaur game on one HTML page.
// The page is a scene of absolute boxes; the frame tick moves the
// display-list operations that make the dinosaur, the obstacles, and the
// ground marks. A key handler never draws, so the game paints from the
// tick alone.
package dino

import (
	"math/rand/v2"
	"time"

	"github.com/chinmay-sawant/go-gpui"
)

// DefaultWidth and DefaultHeight are the size of a newly opened window.
const (
	DefaultWidth  = 900
	DefaultHeight = 300

	// MinWidth and MinHeight are the smallest frame the screen will draw.
	MinWidth  = 540
	MinHeight = 180

	// MaxWidth and MaxHeight cap a resized frame.
	MaxWidth  = 1800
	MaxHeight = 600
)

// point is a scene position in CSS pixels.
type point struct {
	x, y float64
}

// App is the dinosaur game screen.
type App struct {
	page *gpui.Page
	game game
	rng  *rand.Rand

	now  func() time.Time
	last time.Time
	fps  meter

	clouds  [2]point
	pebbles [6]float64

	parts parts
	bound uint64
}
