package page

import "sort"

func keepSpans(source string, spans []controlSpan) []controlSpan {
	ordered := make([]controlSpan, len(spans))
	copy(ordered, spans)
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].Start < ordered[j].Start
	})

	kept := make([]controlSpan, 0, len(ordered))
	prev := 0
	n := len(source)
	for _, sp := range ordered {
		if sp.Start < prev || sp.End > n || sp.End <= sp.Start {
			continue
		}
		kept = append(kept, sp)
		prev = sp.End
	}

	return kept
}
