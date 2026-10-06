package ui

// inList reports whether v is in list.
func inList(list []string, v string) bool {
	for _, s := range list {
		if s == v {
			return true
		}
	}

	return false
}

// sevRank orders the severity names the way the store does.
func sevRank(sev string) int {
	switch sevClass(sev) {
	case "trace":
		return 1
	case "debug":
		return 2
	case "info":
		return 3
	case "warn":
		return 4
	case "error":
		return 5
	case "fatal":
		return 6
	}

	return 0
}

// matched returns the matching entries in ascending ID order.
func (f *fakeFeed) matched(q Query) []Entry {
	var out []Entry

	for _, e := range f.entries {
		if f.match(e, q) {
			out = append(out, e)
		}
	}

	return out
}
