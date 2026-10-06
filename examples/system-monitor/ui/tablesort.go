package ui

import (
	"cmp"
	"sort"
	"strconv"
	"strings"
)

// filtered returns the frozen snapshot's rows matching the query, sorted.
// Filtering and sorting run on an explicit refresh, sort, or filter change,
// never on a per-frame path.
func (t *table) filtered() []Process {
	out := make([]Process, 0, len(t.shown.Procs))

	for _, p := range t.shown.Procs {
		if match(p, t.query) {
			out = append(out, p)
		}
	}

	sort.Slice(out, func(i, j int) bool {
		return t.less(out[i], out[j])
	})

	return out
}

// match compares a process to a case-insensitive name, user, or PID query.
func match(p Process, q string) bool {
	if q == "" {
		return true
	}

	q = strings.ToLower(strings.TrimSpace(q))

	if strings.Contains(strings.ToLower(p.Name), q) {
		return true
	}

	if strings.Contains(strings.ToLower(p.User), q) {
		return true
	}

	return strings.Contains(strconv.Itoa(p.PID), q)
}

// less orders rows by the active key, then by identity, so a tie is stable
// across snapshots however the heavier sort key moves.
func (t *table) less(a, b Process) bool {
	var c int

	switch t.key {
	case sortMem:
		c = cmp.Compare(a.Mem, b.Mem)
	case sortPID:
		c = cmp.Compare(a.PID, b.PID)
	case sortName:
		c = strings.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	default:
		c = cmp.Compare(a.CPU, b.CPU)
	}

	if c != 0 {
		if t.desc {
			return c > 0
		}

		return c < 0
	}

	return a.ID < b.ID
}
