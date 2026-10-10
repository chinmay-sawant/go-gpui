package host

// TouchScrollConfig lets a page tune finger-following and inertial scroll.
type TouchScrollConfig interface {
	TouchScrollSensitivity() float64
	TouchScrollDeceleration() float64
}
