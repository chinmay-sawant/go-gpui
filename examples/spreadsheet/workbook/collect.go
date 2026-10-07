package workbook

// collect finds the affected formula cells with a breadth-first walk over
// dependents, and separately reports the cells past the graph depth limit.
func (w *Workbook) collect(s *Sheet, changed []Pos) (map[Pos]bool, map[Pos]bool) {
	affected := map[Pos]bool{}
	blocked := map[Pos]bool{}
	depth := map[Pos]int{}

	frontier := make([]Pos, 0, len(changed))

	for _, p := range changed {
		if !p.Valid() {
			continue
		}

		depth[p] = 0
		frontier = append(frontier, p)

		if _, isFormula := s.parsed[p]; isFormula {
			affected[p] = true
		}
	}

	for len(frontier) > 0 {
		p := frontier[0]
		frontier = frontier[1:]

		for _, f := range s.dependents(p) {
			if affected[f] {
				continue
			}

			if depth[p]+1 > w.lim.MaxDepth {
				affected[f] = true
				blocked[f] = true

				continue
			}

			affected[f] = true
			depth[f] = depth[p] + 1
			frontier = append(frontier, f)
		}
	}

	return affected, blocked
}
