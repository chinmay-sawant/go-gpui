package workbook

// topo orders the affected formulas so every formula runs after the ones
// it reads. Formulas left over are in a dependency cycle.
func topo(s *Sheet, list []Pos) ([]Pos, []Pos) {
	in := map[Pos]bool{}

	for _, p := range list {
		in[p] = true
	}

	indeg := map[Pos]int{}
	edges := map[Pos][]Pos{}

	for _, g := range list {
		for _, r := range s.parsed[g].refs {
			for _, f := range list {
				if !in[f] || !r.Contains(f.Row, f.Col) {
					continue
				}

				indeg[g]++
				edges[f] = append(edges[f], g)

				break
			}
		}
	}

	zeros := make([]Pos, 0, len(list))

	for _, p := range list {
		if indeg[p] == 0 {
			zeros = append(zeros, p)
		}
	}

	order := make([]Pos, 0, len(list))

	for len(zeros) > 0 {
		f := zeros[0]
		zeros = zeros[1:]
		order = append(order, f)

		for _, g := range edges[f] {
			indeg[g]--

			if indeg[g] == 0 {
				zeros = append(zeros, g)
			}
		}
	}

	done := map[Pos]bool{}
	for _, p := range order {
		done[p] = true
	}

	var cyclic []Pos

	for _, p := range list {
		if !done[p] {
			cyclic = append(cyclic, p)
		}
	}

	return order, cyclic
}
