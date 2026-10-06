package page

// SetPerf turns pipeline timing, dirty counters, and alloc tracking on or
// off. It is off by default, so a shipped app pays nothing; a developer
// opts in to fill the DevTools performance rows and the matching Stats.
func (p *Page) SetPerf(on bool) { p.perf = on }

// Perf reports whether pipeline timing and dirty counters are recorded.
func (p *Page) Perf() bool { return p.perf }

// clearPerf zeroes the opt-in Stats fields after an untracked Redraw, so a
// page that opts out never shows stale developer numbers.
func (p *Page) clearPerf() {
	p.stats.lastTemplate = 0
	p.stats.layoutTime = 0
	p.stats.displayListTime = 0
	p.stats.paintTime = 0
	p.stats.dirtyOps = 0
	p.stats.dirtyRegions = 0
	p.stats.changedOps = 0
	p.stats.allocFrame = 0
}
