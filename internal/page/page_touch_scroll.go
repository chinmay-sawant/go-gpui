package page

import "math"

// TouchScrollOptions tunes touch scrolling. Sensitivity defaults to 1 and is
// limited to 0.25-3. Deceleration defaults to 5; larger values stop sooner.
type TouchScrollOptions struct {
	Sensitivity  float64
	Deceleration float64
}

// SetTouchScrollOptions changes touch scrolling for this page.
func (p *Page) SetTouchScrollOptions(options TouchScrollOptions) {
	if p == nil {
		return
	}
	if options.Sensitivity <= 0 || mathInvalid(options.Sensitivity) {
		options.Sensitivity = 1
	}
	if options.Deceleration <= 0 || mathInvalid(options.Deceleration) {
		options.Deceleration = 5
	}
	p.touchScroll = TouchScrollOptions{
		Sensitivity:  clampPage(options.Sensitivity, .25, 3),
		Deceleration: clampPage(options.Deceleration, 1, 12),
	}
}

// TouchScrollSensitivity is the configured finger movement multiplier.
func (p *Page) TouchScrollSensitivity() float64 {
	if p == nil || p.touchScroll.Sensitivity == 0 {
		return 1
	}
	return p.touchScroll.Sensitivity
}

// TouchScrollDeceleration is the configured fling decay per second.
func (p *Page) TouchScrollDeceleration() float64 {
	if p == nil || p.touchScroll.Deceleration == 0 {
		return 5
	}
	return p.touchScroll.Deceleration
}

// TouchScrollOptions returns the active touch scroll settings.
func (p *Page) TouchScrollOptions() TouchScrollOptions {
	return TouchScrollOptions{Sensitivity: p.TouchScrollSensitivity(), Deceleration: p.TouchScrollDeceleration()}
}

func clampPage(v, low, high float64) float64 { return max(low, min(v, high)) }

func mathInvalid(v float64) bool { return math.IsNaN(v) || math.IsInf(v, 0) }
