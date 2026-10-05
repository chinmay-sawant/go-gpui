package activity

// clone copies the sample feed so two Defaults never share state.
func clone(items []Item) []Item {
	out := make([]Item, len(items))
	copy(out, items)

	return out
}
