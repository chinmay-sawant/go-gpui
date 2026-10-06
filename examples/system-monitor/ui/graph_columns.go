package ui

// sample returns the s-th oldest sample.
func (g *Graph) sample(s int) (float64, bool) {
	i := (g.head - g.n + s + len(g.vals)) % len(g.vals)

	return g.vals[i], g.ok[i]
}

// Columns downsamples the ring to exactly cols columns, oldest first. Each
// column keeps its bucket's maximum so a short spike survives; a column with
// no known samples is a gap. When fewer samples than columns are stored, the
// newest sample lands in the last column and the leading columns are gaps.
func (g *Graph) Columns(cols int) ([]float64, []bool) {
	vals := make([]float64, cols)
	oks := make([]bool, cols)

	if cols < 1 || g.n == 0 {
		return vals, oks
	}

	if g.n < cols {
		off := cols - g.n

		for s := 0; s < g.n; s++ {
			vals[off+s], oks[off+s] = g.sample(s)
		}

		return vals, oks
	}

	for c := 0; c < cols; c++ {
		start, end := c*g.n/cols, (c+1)*g.n/cols
		if end <= start {
			end = start + 1
		}

		m, ok := 0.0, false

		for s := start; s < end; s++ {
			if v, known := g.sample(s); known && (!ok || v > m) {
				m, ok = v, true
			}
		}

		vals[c], oks[c] = m, ok
	}

	return vals, oks
}
