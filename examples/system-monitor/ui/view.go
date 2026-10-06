package ui

// View is the template data. buildView fills every string, so the template
// is markup and ranges only.
type View struct {
	Dark    bool
	Live    bool
	Nav     string
	Mode    string
	At      string
	TableAt string
	Notice  string
	Panels  []Panel
	Query   string
	SortKey string
	SortDir string
	Rows    []Row
	PageNo  int
	Pages   int
	Total   int
	Shown   int
	HasPrev bool
	HasNext bool
	HasNew  bool
	NewText string
	Sel     SelView
}

// Panel is one overview card over a fixed-size graph.
type Panel struct {
	ID    string
	Name  string
	Value string
	Sub   string
	Peak  string
	Bars  []int
	Empty bool
}

// Row is one process table row.
type Row struct {
	ID    string
	PID   string
	Name  string
	User  string
	State string
	CPU   string
	Mem   string
	Sel   bool
}

// SelView is the process detail card.
type SelView struct {
	Any     bool
	ID      string
	Name    string
	PID     string
	Status  string
	Loading bool
	Err     string
	Fields  []Field
}

// Field is one labelled row of the detail card. Key names the paint handle
// the tick rewrites in place, so the row set stays fixed.
type Field struct {
	Key   string
	Label string
	Value string
}
