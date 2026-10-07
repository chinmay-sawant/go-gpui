package domain

import "time"

// Point is one value in a graph buffer. Gap marks a break, so the UI draws the
// line in segments instead of connecting across a suspend or a mode switch.
type Point struct {
	At    time.Time
	Value float64
	Valid bool
	Gap   bool
}

// Ring is a fixed-capacity history of points. Push overwrites the oldest
// point once full, so a graph buffer never grows. Use NewRing; the zero value
// has no capacity.
type Ring struct {
	points []Point
	next   int
	full   bool
}

// NewRing returns a ring that holds capacity points.
func NewRing(capacity int) *Ring {
	if capacity < 1 {
		capacity = 1
	}

	return &Ring{points: make([]Point, capacity)}
}

// Push adds one point, dropping the oldest when the ring is full.
func (r *Ring) Push(p Point) {
	r.points[r.next] = p
	r.next = (r.next + 1) % len(r.points)
	if r.next == 0 {
		r.full = true
	}
}

// Len returns how many points are stored.
func (r *Ring) Len() int {
	if r.full {
		return len(r.points)
	}

	return r.next
}

// Cap returns the fixed capacity.
func (r *Ring) Cap() int { return len(r.points) }

// Reset drops every point and keeps the capacity.
func (r *Ring) Reset() {
	r.next, r.full = 0, false
	clear(r.points)
}

// Points returns a copy of the stored points, oldest first.
func (r *Ring) Points() []Point {
	n := r.Len()
	if n == 0 {
		return nil
	}

	out := make([]Point, 0, n)
	if r.full {
		out = append(out, r.points[r.next:]...)
	}
	out = append(out, r.points[:r.next]...)

	return out
}

// SeriesKey names one graph line: a metric plus the device it belongs to.
type SeriesKey struct {
	Metric Metric
	Device string
}

// Series is one named graph line with its points, oldest first. Points is a
// copy, so the caller may keep it while the collector pushes more.
type Series struct {
	Key    SeriesKey
	Points []Point
}
