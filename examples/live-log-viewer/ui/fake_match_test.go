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
