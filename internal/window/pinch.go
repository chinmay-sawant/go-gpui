package window

import "math"

const (
	minZoom = 0.25
	maxZoom = 4
)

// pinchScale returns the zoom for a pinch: the base zoom times the change in
// finger span, clamped to a usable range.
func pinchScale(start, now, base float64) float64 {
	if start <= 0 || base <= 0 {
		return 1
	}

	z := base * now / start
	if z < minZoom {
		return minZoom
	}

	if z > maxZoom {
		return maxZoom
	}

	return z
}

// touchSpan is the distance between two fingers.
func touchSpan(a, b touchFinger) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)

	return math.Hypot(dx, dy)
}
