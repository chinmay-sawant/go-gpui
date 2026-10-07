package ui

import "time"

// reset clears graph state and snapshots after a mode switch, so the new
// source's baselines never mix with the old source's history. The generation
// advance makes every in-flight result obsolete.
func (s *state) reset() {
	for _, p := range s.panels {
		p.graph.Reset()
		p.last = Reading{}
		p.have = false
	}

	s.at = time.Time{}
	s.table = newTable()
	s.sel = selection{}
	s.haveAny = false
	s.problems = nil
	s.dirtyText = true

	s.gen.Add(1)
}

// apply pushes one reading into its panel. A gap sample records a hole
// instead of a value, so the graph breaks there.
func (s *state) apply(r Reading, gap bool) {
	p, ok := s.panels[r.ID]
	if !ok {
		return
	}

	p.graph.Push(r.Value, r.OK && !gap)
	p.last = r
	p.have = true
	s.haveAny = true
}
