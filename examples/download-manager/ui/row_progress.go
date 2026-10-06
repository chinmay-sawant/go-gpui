package ui

import "strconv"

// Fraction returns Done over Total in 0..1, or -1 when it is unknown.
func (r Row) Fraction() float64 {
	if r.Total <= 0 || r.Done < 0 {
		return -1
	}

	if r.Done >= r.Total {
		return 1
	}

	return float64(r.Done) / float64(r.Total)
}

// Progress returns the percentage, or "--" when the total is unknown.
func (r Row) Progress() string {
	f := r.Fraction()
	if f < 0 {
		return "--"
	}

	return strconv.Itoa(int(f*100+0.5)) + "%"
}

// BarWidth is the initial CSS width of the progress fill. The small floor
// keeps the fill operation in the display list while the fraction is zero,
// so the tick can grow it again later.
func (r Row) BarWidth() string {
	f := r.Fraction()
	if f < 0.005 {
		return "0.5%"
	}

	return strconv.FormatFloat(f*100, 'f', 1, 64) + "%"
}
