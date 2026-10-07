package ui

// Graph is a fixed-size ring of metric samples. Push overwrites the oldest
// sample once the ring is full, so memory stays bounded by the capacity.
type Graph struct {
	vals []float64
	ok   []bool
	head int
	n    int
}

// NewGraph returns a ring that holds capacity samples.
func NewGraph(capacity int) *Graph {
	if capacity < 1 {
		capacity = 1
	}

	return &Graph{vals: make([]float64, capacity), ok: make([]bool, capacity)}
}

// Push appends one sample, dropping the oldest when full. ok=false stores an
// explicit gap, such as a counter reset or an unavailable sensor.
func (g *Graph) Push(v float64, ok bool) {
	g.vals[g.head] = v
	g.ok[g.head] = ok
	g.head = (g.head + 1) % len(g.vals)

	if g.n < len(g.vals) {
		g.n++
	}
}

// Reset clears every sample, for a collector or mode switch.
func (g *Graph) Reset() {
	g.head, g.n = 0, 0

	for i := range g.ok {
		g.ok[i] = false
	}
}

// Len returns the number of stored samples.
func (g *Graph) Len() int {
	return g.n
}

// Last returns the newest sample.
func (g *Graph) Last() (float64, bool) {
	if g.n == 0 {
		return 0, false
	}

	i := (g.head - 1 + len(g.vals)) % len(g.vals)

	return g.vals[i], g.ok[i]
}

// Max returns the largest known sample.
func (g *Graph) Max() (float64, bool) {
	m, ok := 0.0, false

	for s := 0; s < g.n; s++ {
		if v, known := g.sample(s); known && (!ok || v > m) {
			m, ok = v, true
		}
	}

	return m, ok
}
