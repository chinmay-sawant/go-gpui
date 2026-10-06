package window

import (
	"sort"
	"time"
)

const perfRingCap = 120

const perfLongThreshold = 25 * time.Millisecond

// perfRing keeps the last perfRingCap frame durations.
type perfRing struct {
	buf  [perfRingCap]time.Duration
	n    int
	next int
}

func (r *perfRing) record(d time.Duration) {
	if d < 0 {
		d = 0
	}
	r.buf[r.next] = d
	r.next = (r.next + 1) % perfRingCap
	if r.n < perfRingCap {
		r.n++
	}
}

func (r *perfRing) count() int { return r.n }

func (r *perfRing) copy() []time.Duration {
	out := make([]time.Duration, r.n)
	old := (r.next - r.n + perfRingCap) % perfRingCap
	for i := range out {
		out[i] = r.buf[(old+i)%perfRingCap]
	}
	return out
}

func (r *perfRing) avg() time.Duration {
	if r.n == 0 {
		return 0
	}
	var sum time.Duration
	for _, d := range r.copy() {
		sum += d
	}
	return sum / time.Duration(r.n)
}

func (r *perfRing) quantile(q float64) time.Duration {
	if r.n == 0 {
		return 0
	}
	s := r.copy()
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	at := q * float64(r.n-1)
	lo := int(at)
	if lo >= r.n-1 {
		return s[r.n-1]
	}
	fr := at - float64(lo)
	return s[lo] + time.Duration(fr*float64(s[lo+1]-s[lo]))
}

func (r *perfRing) p95() time.Duration { return r.quantile(0.95) }

func (r *perfRing) p99() time.Duration { return r.quantile(0.99) }

func (r *perfRing) longFrames(th time.Duration) int {
	c := 0
	for _, d := range r.copy() {
		if d > th {
			c++
		}
	}
	return c
}
