package ui

import "fmt"

// sortArrow marks the active sort direction.
func sortArrow(desc bool) string {
	if desc {
		return "\u2193"
	}

	return "\u2191"
}

// newText labels the refresh indicator.
func newText(hasNew bool) string {
	if hasNew {
		return "new sample ready"
	}

	return ""
}

// modeLine is the sidebar status. The tick rewrites it in place; the full
// collector problem appears in the notice on the next redraw.
func modeLine(s *state) string {
	switch n := len(s.problems); {
	case n == 1:
		return s.mode + " mode \u00b7 1 issue"
	case n > 1:
		return fmt.Sprintf("%s mode \u00b7 %d issues", s.mode, n)
	default:
		return s.mode + " mode"
	}
}
