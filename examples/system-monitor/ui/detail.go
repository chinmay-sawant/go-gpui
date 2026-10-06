package ui

// detailRows is the fixed detail-card row order. Every row always renders,
// so the tick rewrites values in place without a Redraw.
var detailRows = []struct{ key, label string }{
	{"status", "Status"},
	{"state", "State"},
	{"cpu", "CPU"},
	{"mem", "Memory"},
	{"threads", "Threads"},
	{"handles", "Handles"},
	{"priority", "Priority"},
	{"started", "Started"},
	{"virtual", "Virtual"},
	{"io", "Disk I/O"},
	{"parent", "Parent"},
	{"exe", "Executable"},
	{"cwd", "Working directory"},
	{"command", "Command line"},
}

// detailFields renders the fixed rows for the selection. A denied or absent
// value reads n/a instead of a zero.
func detailFields(sel selection) []Field {
	fields := make([]Field, 0, len(detailRows))

	for _, r := range detailRows {
		fields = append(fields, Field{
			Key:   r.key,
			Label: r.label,
			Value: detailValue(r.key, sel),
		})
	}

	return fields
}
