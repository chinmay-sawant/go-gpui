package ui

import (
	"errors"
	"strings"
)

func (b *fakeBackend) Range(id string, a Area) ([]Cell, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if a.Empty() {
		return nil, nil
	}

	b.fetches = append(b.fetches, a)
	out := make([]Cell, 0, a.Count())
	for r := a.R0; r <= a.R1; r++ {
		for c := a.C0; c <= a.C1; c++ {
			out = append(out, b.cells[id][[2]int{r, c}])
		}
	}

	return out, nil
}

func (b *fakeBackend) Used(id string) (Area, bool) {
	b.mu.Lock()
	defer b.mu.Unlock()

	cells := b.cells[id]
	if len(cells) == 0 {
		return Area{}, false
	}

	var out Area
	first := true
	for k := range cells {
		if first {
			out = Area{k[0], k[1], k[0], k[1]}
			first = false

			continue
		}

		out.R0, out.C0 = min(out.R0, k[0]), min(out.C0, k[1])
		out.R1, out.C1 = max(out.R1, k[0]), max(out.C1, k[1])
	}

	return out, true
}

func (b *fakeBackend) Apply(id string, edits []Edit) (uint64, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.failSet {
		return b.rev, errors.New("disk full")
	}

	b.applied = append(b.applied, edits)
	b.rev++
	if b.cells[id] == nil {
		b.cells[id] = map[[2]int]Cell{}
	}

	for _, ed := range edits {
		key := [2]int{ed.Row, ed.Col}
		cell := Cell{Raw: ed.Raw, Display: ed.Raw}
		if strings.HasPrefix(ed.Raw, "=") {
			cell.Display = "=?"
			cell.Num = true
		}

		if ed.Raw == "" {
			delete(b.cells[id], key)

			continue
		}

		b.cells[id][key] = cell
	}

	return b.rev, nil
}
