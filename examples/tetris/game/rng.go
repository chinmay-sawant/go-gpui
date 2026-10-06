package game

// rng is a splitmix64 generator. Its whole state is one uint64, so a
// snapshot can save and restore the piece sequence exactly.
type rng struct{ state uint64 }

// next returns the next 64 random bits.
func (r *rng) next() uint64 {
	r.state += 0x9e3779b97f4a7c15
	z := r.state
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb

	return z ^ (z >> 31)
}

// intn returns a value in [0, n).
func (r *rng) intn(n int) int { return int(r.next() % uint64(n)) }
