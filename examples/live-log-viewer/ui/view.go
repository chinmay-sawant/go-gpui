package ui

// Row is one fixed-height summary row the list renders.
type Row struct {
	ID        int64
	Time      string
	Source    string
	Severity  string
	SevClass  string
	Text      string
	Lines     int
	Truncated bool
	Partial   bool
	Lost      bool
	Selected  bool
	Alt       bool
}

// Chip is one severity filter toggle.
type Chip struct {
	Name  string
	Label string
	Class string
	On    bool
}

// DetailView is one entry opened in the detail pane.
type DetailView struct {
	Meta      string
	Lines     []string
	Truncated bool
	Partial   bool
}

// View is the whole template data.
type View struct {
	Mode       string
	Rows       []Row
	Sources    []SourceRow
	Severities []Chip
	Detail     DetailView

	QueryText string
	TopPad    int
	BotPad    int
	BarTop    int
	SideH     int
	StatusTop int

	Status    string
	Right     string
	PageLabel string
	Empty     string
	Follow    bool
	Paused    bool
	Dark      bool
	Unread    int
	HasOlder  bool
	HasNewer  bool
}

// sevNames lists the severity filter chips in display order.
var sevNames = [...][2]string{
	{"trace", "TRACE"}, {"debug", "DEBUG"}, {"info", "INFO"},
	{"warn", "WARN"}, {"error", "ERROR"}, {"fatal", "FATAL"},
}

// chips builds the severity toggles for the active set.
func chips(active []string) []Chip {
	on := map[string]bool{}
	for _, s := range active {
		on[sevClass(s)] = true
	}

	out := make([]Chip, 0, len(sevNames))
	for _, pair := range sevNames {
		out = append(out, Chip{
			Name: pair[0], Label: pair[1], Class: pair[0], On: on[pair[0]],
		})
	}

	return out
}
